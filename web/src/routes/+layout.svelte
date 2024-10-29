<script lang="ts">
	import { SvelteUIProvider, AppShell, Navbar, Header } from '@svelteuidev/core';
	import HeadContent from './HeadContent.svelte';
	import NavContent from './NavContent.svelte';

	import type { LayoutData } from './$types';
	
	export let data: LayoutData;
</script>

<SvelteUIProvider withGlobalStyles themeObserver={'dark'}>
	<AppShell>
		<Navbar slot="navbar" position={{ top: 90, left: 15 }} width={{base: data.session == null ? 0 : 150, sm: data.session == null ? 0 : 150, lg: data.session == null ? 0 : 150}} hidden={data.session == null}>
			<NavContent />
		</Navbar>
		<Header slot="header" fixed height=75 style="padding: 15px 15px 0 15px;">
			<HeadContent />
		</Header>
		<div slot="aside" id="aside" style="position: relative; top: 110px; right: 20px; max-height: calc(100vh - 120px);" />

		<slot />
	</AppShell>
</SvelteUIProvider>

<style>
	:global(body) {
		margin: 0;
	}
</style>
