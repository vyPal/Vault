<script lang="ts">
	import FileBrowser from '$lib/components/FileBrowser.svelte';
	import type { PageData } from './$types';
	export let data: PageData;
</script>

<h1>Files	</h1>

<FileBrowser basePath={data.session?.username + '/vault/'} />
