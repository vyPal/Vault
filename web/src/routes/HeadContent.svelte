<script lang="ts">
   import { signOut } from "@auth/sveltekit/client"
   import { Menu } from "@svelteuidev/core";
   import { page } from "$app/stores"
   import Avatar from '@svelte-put/avatar/Avatar.svelte';
   import { SignIn } from "@auth/sveltekit/components";
</script>

<a href="/">
<img src="Mimlex.png" alt="Mimlex logo" width="50" style="display: inline-block;"/>
<h1 style="display: inline-block; position: relative; top: -15px; color: white;">Mimlex <span style="background: -webkit-linear-gradient(#ddd, #333); -webkit-background-clip: text; -webkit-text-fill-color: transparent;">Vault</span></h1></a>

<div style="float: right;">
   <Menu>
      <Avatar size={50} uiAvatar={$page.data.session?.user?.name ?? undefined} slot="control" style="border-radius: 50%;"/>
      {#if $page.data.session}
         <Menu.Label>{$page.data.session?.user?.name}</Menu.Label>
         <Menu.Item on:click={() => signOut()}>Sign Out</Menu.Item>
      {:else}
         <SignIn provider="authentik" className="signinform" signInPage="signin">
            <div slot="submitButton">
               <Menu.Item>Sign In</Menu.Item>
				    </div>
				</SignIn>
      {/if}
   </Menu>
</div>
