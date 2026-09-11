<script lang="ts">
  import IconBase from './IconBase.svelte';
  let props: { size?: number; stroke?: number; color?: string; class?: string; style?: string } =
    $props();
</script>

<IconBase {...props}>
  <path
    d="M9 17v-2m3 2v-4m3 4v-6m2 10H7a2 2 0 01-2-2V5a2 2 0 012-2h5.586a1 1 0 01.707.293l5.414 5.414a1 1 0 01.293.707V19a2 2 0 01-2 2z"
  />
</IconBase>
