<script lang="ts">
  import type { Snippet } from 'svelte';

  let {
    size = 24,
    stroke = 2,
    color = 'currentColor',
    class: className = '',
    style = '',
    children,
  }: {
    size?: number;
    stroke?: number;
    color?: string;
    class?: string;
    style?: string;
    children?: Snippet;
  } = $props();
</script>

<svg
  xmlns="http://www.w3.org/2000/svg"
  class={className}
  {style}
  width={size}
  height={size}
  viewBox="0 0 24 24"
  fill="none"
  stroke={color}
  stroke-width={stroke}
  stroke-linecap="round"
  stroke-linejoin="round"
>
  {#if children}
    {@render children()}
  {/if}
</svg>
