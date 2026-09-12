<script lang="ts">
  import IconBase from './IconBase.svelte';
  let props: { size?: number; stroke?: number; color?: string; class?: string; style?: string } =
    $props();
</script>

<IconBase {...props}>
  <polyline points="20 6 9 17 4 12" />
</IconBase>
