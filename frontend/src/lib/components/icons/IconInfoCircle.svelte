<script lang="ts">
  import IconBase from './IconBase.svelte';
  let props: { size?: number; stroke?: number; color?: string; class?: string; style?: string } =
    $props();
</script>

<IconBase {...props}>
  <path d="M13 16h-1v-4h-1m1-4h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z" />
</IconBase>
