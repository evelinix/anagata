<script lang="ts">
  import IconBase from './IconBase.svelte';
  let props: { size?: number; stroke?: number; color?: string; class?: string; style?: string } =
    $props();
</script>

<IconBase {...props}>
  <path
    d="M21 12a9 9 0 01-9 9m9-9a9 9 0 00-9-9m9 9H3m9 9a9 9 0 01-9-9m9 9c1.657 0 3-4.03 3-9s-1.343-9-3-9m0 18c-1.657 0-3-4.03-3-9s1.343-9 3-9m-9 9a9 9 0 019-9"
  />
</IconBase>
