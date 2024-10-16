<script>
	import { SvelteUIProvider, AppShell, Navbar, Header } from '@svelteuidev/core';
	import HeadContent from './HeadContent.svelte';
	import NavContent from './NavContent.svelte';
</script>

<SvelteUIProvider withGlobalStyles themeObserver={'dark'}>
	<AppShell>
		<Navbar slot="navbar"
			width={{
			sm: 200,

			lg: 250,

			base: 100
			}}
		>
			<NavContent />
		</Navbar>
		<Header slot="header" height=75>
			<HeadContent />
		</Header>

		<slot />
	</AppShell>
</SvelteUIProvider>
