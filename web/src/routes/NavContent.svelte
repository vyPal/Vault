<script lang="ts">
  import { Button, Stack, Paper, Progress, Text } from '@svelteuidev/core';
  import Icon from '@iconify/svelte';
  import { onMount } from 'svelte';
  import { page } from "$app/stores"

  let usage = 0;
  onMount(async () => {
    const response = await fetch(`/files/usage/${$page.data.session?.username}`);
    const data = await response.json();
    usage = parseInt(data.size);
  });

  function formatBytes(bytes: number, decimals = 1) {
    if (bytes === 0) return '0 Bytes';
    
    const k = 1024;
    const sizes = ['Bytes', 'KB', 'MB', 'GB', 'TB', 'PB', 'EB', 'ZB', 'YB'];
    const i = Math.floor(Math.log(bytes) / Math.log(k));

    const formattedSize = parseFloat((bytes / Math.pow(k, i)).toFixed(decimals));
    return `${formattedSize} ${sizes[i]}`;
  }
</script>

<Stack justify="space-between">
  <Stack align="strech" justify="flex-start" style="margin-top: 15px; padding-right: 18px;">
    <Button variant="default" style="width: 100%; justify-content: start;" radius={10} href="/home"><Icon icon="material-symbols:home-outline-rounded" /> &nbsp;&nbsp; Home</Button>
    <Button variant="default" style="width: 100%; justify-content: start;" radius={10} href="/apps"><Icon icon="tabler:apps" width="18" height="18" /> &nbsp;&nbsp; Apps</Button>
    <Button variant="default" style="width: 100%; justify-content: start;" radius={10} href="/files"><Icon icon="ci:files" width="18" height="18" /> &nbsp;&nbsp; Files</Button>
  </Stack>
  <Paper shadow="xl" radius="lg" style="background: var(--svelteui-colors-dark900); margin-right: 18px; position: absolute; bottom: 40px;">
    <Text size="sm" color="gray">{formatBytes(usage)} / {formatBytes(1024*1024*1024*5)}</Text>
    <Progress value={usage / (1024*1024*1024*5) * 100} />
  </Paper>
</Stack>
