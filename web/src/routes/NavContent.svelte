<script>
  import { Button, Stack } from '@svelteuidev/core';
  import Icon from '@iconify/svelte';
</script>

<Stack align="strech" justify="flex-start" style="margin-top: 15px;">
  <Button variant="default" radius={30}><Icon icon="material-symbols:home-outline-rounded" /> &nbsp;&nbsp; Home</Button>
  <Button variant="default" radius={30}>2</Button>
  <Button variant="default" radius={30}>3</Button>
</Stack>
