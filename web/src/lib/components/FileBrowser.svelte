<script lang="ts">
  import { onMount } from "svelte";
  import { Alert, Text, Grid, ThemeIcon, Stack, Group, Paper, Breadcrumbs, Space, Button, Flex, NativeSelect, ChipGroup, Menu, Modal, Affix, Notification, Progress } from "@svelteuidev/core";
  import Icon from "@iconify/svelte";
  import DropFile from 'svelte-parts/DropFile.svelte'

  export let basePath: string = "/";

  let currentPath = basePath;

  type MyFile = {
    isFolder: boolean;
    name: string;
    extension: string;
    fullPath: string;
    icon: string;
    size: number;
    lastModified: Date;
    previewHtml: DynComponent;
  };

  export let files: MyFile[] = [];

  const getIcon = (ext: string) => {
    switch (ext) {
      case "pdf":
        return "ph:file-pdf";
      case "doc":
      case "docx":
        return "ph:file-doc";
      case "xls":
      case "xlsx":
        return "ph:file-xls";
      case "ppt":
      case "pptx":
        return "ph:file-ppt";
      case "zip":
        return "ph:file-zip";
      case "rar":
        return "ph:file-archive";
      case "jpg":
      case "jpeg":
        return "ph:file-jpg";
      case "png":
        return "ph:file-png";
      case "gif":
        return "ph:file-image";
      case "svg":
        return "ph:file-svg";
      case "mp3":
      case "wav":
      case "ogg":
        return "ph:file-audio";
      case "mp4":
      case "avi":
      case "mov":
      case "wmv":
        return "ph:file-video";
      default:
        return "ph:file";
    }
  }

  const getPreview = (name: string, ext: string) => {
    switch (ext) {
      case "pdf":
        return {
          type: "html",
          html: `<iframe src="/files/get/${name}" style="width: 100%; height: 100%; border: none;" title="fileModalSrc"></iframe>`
        }
      case "jpg":
      case "jpeg":
      case "png":
      case "gif":
      case "svg":
        return {
          type: "html",
          html: `<img src="/files/get/${name}" style="height: 100%;" />`
        }
      case "mp3":
      case "wav":
      case "ogg":
        return {
          type: "html",
          html: `<audio src="/files/get/${name}" controls style="width: 100%; height: 100%;"></audio>`
        }
      case "mp4":
      case "avi":
      case "mov":
      case "wmv":
        return {
          type: "html",
          html: `<video src="/files/get/${name}" controls style="height: 100%;"></video>`
        }
      default:
        return {
          type: "component",
          component: Icon,
          props: {
            icon: getIcon(ext),
            width: "82",
            height: "82"
          }
        }
    }
  }

  const fetchFiles = async (path: string) => {
    currentPath = path;
    let nf = await fetch(`/files/list/${path}`).then(async (res) => res.json()).catch((e) => console.error(e));
    files = nf.map((file: any) => {
      if (file.name.endsWith("/")) {
        return {
          isFolder: true,
          name: file.name.split("/").slice(-2)[0],
          extension: "",
          fullPath: file.name,
          icon: "ph:folder",
          size: 0,
          lastModified: Date.parse(file.lastModified),
          previewHtml: {
            type: "component",
            component: Icon,
            props: {
              icon: "ph:folder",
              width: "82",
              height: "82"
            }
          }
        };
      }
      return {
        isFolder: false,
        name: file.name.split("/").slice(-1)[0],
        extension: file.name.split(".").pop(),
        fullPath: file.name,
        icon: getIcon(file.name.split(".").pop()),
        size: file.size,
        lastModified: Date.parse(file.lastModified),
        previewHtml: getPreview(file.name, file.name.split(".").pop())
      };
    })
  };

  onMount(() => {
    fetchFiles(basePath);
  });

  type DynComponent = {
    type: "component";
    component: any;
    props: any;
  } | {
    type: "html";
    html: string;
  } | {
    type: "text";
    text: string;
  };

  const openFile = (file: MyFile) => {
    if (file.isFolder) {
      fetchFiles(file.fullPath);
      currentPath = file.fullPath;
    } else {
      // Placeholder
    }
  }

  let sortBy = "Name A-Z";

  $: files = files.sort((a, b) => {
    switch (sortBy) {
      case "Name A-Z":
        return a.name.localeCompare(b.name);
      case "Name Z-A":
        return b.name.localeCompare(a.name);
      case "Size Asc":
        return a.size - b.size;
      case "Size Desc":
        return b.size - a.size;
      case "Date Asc":
        return a.lastModified.getTime() - b.lastModified.getTime();
      case "Date Desc":
        return b.lastModified.getTime() - a.lastModified.getTime();
      default:
        return 0;
    }
  });

  let viewType = "grid";

  let fileModal = false;

  const onDrop = (files: File[]) => {
    uploadFiles(files);
    fileModal = false;
  }

  let uploadProgress: {[index: string]: number} = {};

  async function uploadFiles(files: File[]) {
    for (const file of files) {
      const formData = new FormData();
      formData.append("file", file);

      await uploadFileWithProgress(formData, file.name);
    }
  }

  function uploadFileWithProgress(formData: FormData, filename: string) {
    return new Promise((resolve, reject) => {
      const xhr = new XMLHttpRequest();
      xhr.open("POST", "/files/push/" + currentPath + filename);

      xhr.upload.onprogress = (event) => {
        if (event.lengthComputable) {
          uploadProgress[filename] = (event.loaded / event.total) * 100;
        }
      };

      xhr.onload = () => {
        if (xhr.status === 200) {
          delete uploadProgress[filename];
          resolve("");
        } else {
          reject(xhr.statusText);
        }
      };
      xhr.onerror = () => reject("Upload failed");
      xhr.send(formData);
    });
  }

  function handleDragEnter(event: DragEvent) {
    event.preventDefault();
    fileModal = true;
  }

  function handleDragLeave(event: DragEvent) {
    event.preventDefault();
    if (event.target === document.body) {
      fileModal = false;
    }
  }

  function formatBytes(bytes: number, decimals = 1) {
    if (bytes === 0) return '0 Bytes';
    
    const k = 1024;
    const sizes = ['Bytes', 'KB', 'MB', 'GB', 'TB', 'PB', 'EB', 'ZB', 'YB'];
    const i = Math.floor(Math.log(bytes) / Math.log(k));

    const formattedSize = parseFloat((bytes / Math.pow(k, i)).toFixed(decimals));
    return `${formattedSize} ${sizes[i]}`;
  }
</script>

<svelte:window on:dragenter={handleDragEnter} on:dragleave={handleDragLeave} />

<h1 style="margin-top: 15px;">Files</h1>

<Flex justify="space-between">
  <Group>
    {#if currentPath != basePath}
      <Button variant="subtle" size="xs" on:click={(_) => fetchFiles(currentPath.split("/").slice(0, -2).join("/")+"/")}>
        <Icon icon="ph:arrow-left" width="17" height="17" />
        <Text>Back</Text>
      </Button>
    {/if}
    <Breadcrumbs size="md">
      {@const pathParts = currentPath.substring(basePath.length).split('/').slice(0, -1)}
      <Breadcrumbs.Item active={pathParts.length == 0}>
        <Icon icon="ph:house" width="17" height="17" />
      </Breadcrumbs.Item>
      {#if pathParts.length == 1}
        <Breadcrumbs.Item active={true}>
          <Text>{pathParts[0]}</Text>
        </Breadcrumbs.Item>
      {:else if pathParts.length > 1}
        {#each pathParts.slice(0, pathParts.length - 1) as part}
          <Breadcrumbs.Item>
            <Text>{part}</Text>
          </Breadcrumbs.Item>
        {/each}
        <Breadcrumbs.Item active={true}>
          <Text>{pathParts[pathParts.length - 1]}</Text>
        </Breadcrumbs.Item>
      {/if}
    </Breadcrumbs>
  </Group>
  <Group>
    <Menu>
      <Button slot="control" variant="default">Upload</Button>
      <Menu.Item on:click={(_) => {fileModal = true}}>Upload File</Menu.Item>
      <Menu.Item>Upload Folder</Menu.Item>
    </Menu>
    <NativeSelect data={['Name A-Z', 'Name Z-A', 'Size Asc', 'Size Desc', 'Date Asc', 'Date Desc']} bind:value={sortBy} />
    <ChipGroup variant="filled" bind:value={viewType} items={[{label: "Grid", value: "grid"}, {label: "List", value: "list"}]} />
  </Group>
</Flex>
<Space h={18} />

{#if files.length === 0}
  <Alert>No files found.</Alert>
{:else}
  {#if viewType === "grid"}
    <Grid>
      {#each files as file}
        <button style="background: none; margin: 5px; padding: 0; border: 0;" on:click={(_) => openFile(file)}>
          <Paper shadow="xl" radius="lg" style="background: var(--svelteui-colors-dark900); padding: 15px; max-width: 250px; max-height: 160;">
            <div style="width: 220px; height: 100px; display: flex; justify-content: center;">
              {#if file.previewHtml.type === 'component'}
                {@const DynamicComponent = file.previewHtml.component}
                <DynamicComponent {...file.previewHtml.props} />
              {:else if file.previewHtml.type === 'html'}
                {@html file.previewHtml.html}
              {:else if file.previewHtml.type === 'text'}
                <Text>{file.previewHtml.text}</Text>
              {/if}
            </div>
            <Space h={12} />
            <Group>
              <ThemeIcon color="gray" radius="xl" size="lg">
                <Icon icon={file.icon} width="17" height="17" />
              </ThemeIcon>
              <Text>{file.name}</Text>
            </Group>
          </Paper>
        </button>
      {/each}
    </Grid>
  {:else if viewType === "list"}
    <Stack spacing="xs">
      {#each files as file}
        <button style="background: none; margin: 5px; padding: 0; border: 0;" on:click={(_) => openFile(file)}>
          <Paper shadow="xl" radius="lg" style="background: var(--svelteui-colors-dark900); padding: 15px;">
            <Group>
              <div style="width: 220px; height: 100px; display: flex; justify-content: center;">
                {#if file.previewHtml.type === 'component'}
                  {@const DynamicComponent = file.previewHtml.component}
                  <DynamicComponent {...file.previewHtml.props} />
                {:else if file.previewHtml.type === 'html'}
                  {@html file.previewHtml.html}
                {:else if file.previewHtml.type === 'text'}
                  <Text>{file.previewHtml.text}</Text>
                {/if}
                <Space w={12} />
              </div>
              <Stack>
                <Group>
                  <ThemeIcon color="gray" radius="xl" size="lg">
                    <Icon icon={file.icon} width="17" height="17" />
                  </ThemeIcon>
                  <Text>{file.name}</Text>
                </Group>
                <Space h={12} />
                <Text size="sm" color="gray">{formatBytes(file.size)}</Text>
                <Text size="sm" color="gray">{new Date(file.lastModified).toLocaleString()}</Text>
              </Stack>
            </Group>
          </Paper>
        </button>
      {/each}
    </Stack>
  {/if}
{/if}

<Modal opened={fileModal} on:close={(_) => {fileModal = false}} centered size="85%">
  <DropFile onDrop={onDrop} />  
</Modal>

<Affix position={{ bottom: 20, right: 20 }}>
  {#each Object.keys(uploadProgress) as filename}
    <Notification title="Upload Progress" loading={uploadProgress[filename] != 100} color="green">
      <Stack>
        <Text>{filename}</Text>
        <Progress value={uploadProgress[filename]} />
      </Stack>
    </Notification>
  {/each}
</Affix>
