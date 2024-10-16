<script lang="ts">
   import { signIn } from "@auth/sveltekit/client"
</script>

<img src="Mimlex.png" alt="Mimlex logo" width="50" style="display: inline-block;"/>
<h1 style="display: inline-block; position: relative; top: -15px; ">Mimlex <span style="background: -webkit-linear-gradient(#ddd, #333); -webkit-background-clip: text; -webkit-text-fill-color: transparent;">Vault</span></h1>

<div style="float: right;">
	<button on:click={signIn} style="background: #333; color: #ddd; border: none; padding: 10px 20px; border-radius: 5px; cursor: pointer;">Sign In</button>
</div>
