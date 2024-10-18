<script>
	import Icon from "@iconify/svelte";
	import {Button, Center, Paper} from "@svelteuidev/core";
	import {page} from "$app/stores";
	import { SignIn } from "@auth/sveltekit/components";
</script>

<Center>
	<Paper shadow="xl" radius="xl" style="background: var(--svelteui-colors-dark900); padding: 50px;">
		<h1 style="display: inline-block; position: relative; top: -15px; font-size: 72px;">Mimlex <span style="background: -webkit-linear-gradient(#ddd, #333); -webkit-background-clip: text; -webkit-text-fill-color: transparent;">Vault</span> <Icon icon="material-symbols-light:shield-lock-rounded" width="64" height="64" /></h1>
		<p>Your secure cloud storage</p>
		{#if $page.data.session != null}
			<a href="home"><Button variant="default" style="width: 100%; left: auto; right: auto;">Take me there <Icon icon="material-symbols:keyboard-arrow-right" width="24px"/></Button></a>
		{:else}
			<div class="signinform">
				<SignIn provider="authentik" className="signinform" signInPage="signin">
					<div slot="submitButton">
						<Button variant="default" style="width: 100%; left: auto; right: auto; font-size: 18px;">Sign in with Mimlex<span style="background: linear-gradient(325deg, rgba(53,119,255,1) 0%, rgba(255,34,137,1) 100%); -webkit-background-clip: text; -webkit-text-fill-color: transparent;">ID</span></Button>
					</div>
				</SignIn>
			</div>
		{/if}
	</Paper>
</Center>

<style>
	:global(.signinform > button) {
		padding: 0px;
		border: none;
		background: none;
		width: 100%;
	}
</style>
