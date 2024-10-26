<script lang="ts">
	import { SvelteUIProvider, AppShell, Navbar, Header } from '@svelteuidev/core';
	import HeadContent from './HeadContent.svelte';
	import NavContent from './NavContent.svelte';

	import type { LayoutData } from './$types';
	
	export let data: LayoutData;
</script>

<SvelteUIProvider withGlobalStyles themeObserver={'dark'}>
	<AppShell>
		<Navbar slot="navbar" fixed position={{ top: 90, left: 15 }} width={{base: data.session == null ? 0 : 150, sm: data.session == null ? 0 : 150, lg: data.session == null ? 0 : 150}} hidden={data.session == null}>
			<NavContent />
		</Navbar>
		<Header slot="header" height=75 fixed position={{ top: 15, left: 15, right: 15 }}>
			<HeadContent />
		</Header>

		<slot style="margin-left: 150;"/>
	</AppShell>
</SvelteUIProvider>
