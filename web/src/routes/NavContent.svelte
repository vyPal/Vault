<script>
  import { Button, Stack } from '@svelteuidev/core';
  import Icon from '@iconify/svelte';
</script>

<Stack align="strech" justify="flex-start" style="margin-top: 15px; padding-right: 18px;">
  <Button variant="default" style="width: 100%; justify-content: start;" radius={10}><Icon icon="material-symbols:home-outline-rounded" /> &nbsp;&nbsp; Home</Button>
  <Button variant="default" style="width: 100%; justify-content: start;" radius={10}>2</Button>
  <Button variant="default" style="width: 100%; justify-content: start;" radius={10}>3</Button>
</Stack>
