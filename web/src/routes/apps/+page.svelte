<script lang="ts">
	import { Alert, Paper, Group } from '@svelteuidev/core';
	import FileBrowser from '$lib/components/FileBrowser.svelte';
	import type { PageData } from './$types';
	export let data: PageData;
</script>

<h1>Apps</h1>

{#if data.error}
	<Alert>{data.error}</Alert>
{:else}
	<Group>
		{#each data.clientlist as client}
			<Paper shadow="xl" radius="lg" style="background: var(--svelteui-colors-dark900); padding: 15px; margin-bottom: 20px; width: 100%;">
				<h2 style="margin-top: 8px;">{client}</h2>
				<FileBrowser basePath={data.session?.username + '/' + client + '/'} />
			</Paper>
		{/each}
		{#if data.clientlist.length === 0}
			<Paper shadow="xl" radius="lg" style="background: var(--svelteui-colors-dark900); padding: 15px; margin-bottom: 20px; width: 250px;">
				<h2 style="margin-top: 8px;">No apps found</h2>
				<p>It looks like no apps have requested permission to store files on Mimlex Vault yet.</p>
			</Paper>
		{/if}
	</Group>
{/if}
