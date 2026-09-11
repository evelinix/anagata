<script lang="ts">
  import IconBase from './IconBase.svelte';
  let props: { size?: number; stroke?: number; color?: string; class?: string; style?: string } =
    $props();
</script>

<IconBase {...props}>
  <path
    d="M12 9v2m0 4h.01m-6.938 4h13.856c1.54 0 2.502-1.667 1.732-3L13.732 4c-.77-1.333-2.694-1.333-3.464 0L3.34 16c-.77 1.333.192 3 1.732 3z"
  />
</IconBase>
