<script lang="ts">
	import type { PageData } from './$types';
	export let data: PageData;
</script>

<p>Clients: {data.clientlist}</p><br />
<p>Error: {data.error}</p><br />
