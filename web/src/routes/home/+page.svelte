<script lang="ts">
	import { Alert, Paper, Skeleton, Group, Button, Text, Stack, ThemeIcon } from '@svelteuidev/core';
	import type { PageData } from './$types';
	import Icon from '@iconify/svelte';
	export let data: PageData;

	const isFolder = (path: string) => path.endsWith('/');

	const loadFolder = async (path: string, files: any) => {
		console.log(path);
		console.log("f0",files);
		files = files.concat(await fetch(`/files/list/${path}`).then(async res => res.json()));
		console.log("f1",files);
	}
</script>

<h1 style="margin-top: 15px;">Apps</h1>

{#if data.error}
	<Alert>{data.error}</Alert>
{:else}
	<Group>
		{#each data.clientlist ?? [] as client}
			<Paper shadow="xl" radius="lg" style="background: var(--svelteui-colors-dark900); padding: 15px; margin-bottom: 20px; width: 250px;">
				<h2 style="margin-top: 8px;">{client.name}</h2>
				{#await client.files}
					<Skeleton height={8} radius="xl" />
				{:then files}
					<Stack spacing="xs">
						{#each files as file}
							{#if isFolder(file.name)}
								<Button variant="subtle" on:click={(_) => loadFolder(file.name, files)}>
									<ThemeIcon color="blue" radius="xl" size="lg">
										<Icon icon="material-symbols:folder-outline" width="17" height="17" />
									</ThemeIcon>
									<Text weight={500}>{file.name.split("/").slice(2).join("/")}</Text>
								</Button>
							{:else}
								<Button variant="subtle">
									<ThemeIcon color="gray" radius="xl" size="lg">
										<Icon icon="mdi:file-outline" width="17" height="17" />
									</ThemeIcon>
									<Text>{file.name.split("/").slice(2).join("/")}</Text>
								</Button>
							{/if}
						{/each}
					</Stack>
				{:catch error}
					<Alert>{error.message}</Alert>
				{/await}
			</Paper>
		{/each}
	</Group>
{/if}

