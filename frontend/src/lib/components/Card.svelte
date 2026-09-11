<script lang="ts">
  import type { Snippet } from 'svelte';

  type Variant = 'default' | 'filled' | 'outlined' | 'acrylic';

  let {
    variant = 'default',
    padding = true,
    children,
    header,
    footer,
  }: {
    variant?: Variant;
    padding?: boolean;
    children?: Snippet;
    header?: Snippet;
    footer?: Snippet;
  } = $props();
</script>

<div class="ants-card ants-card--{variant}">
  {#if header}
    <div class="ants-card-header">
      {@render header()}
    </div>
  {/if}
  <div class="ants-card-body" class:ants-card-body--padded={padding}>
    {#if children}
      {@render children()}
    {/if}
  </div>
  {#if footer}
    <div class="ants-card-footer">
      {@render footer()}
    </div>
  {/if}
</div>

<style>
  .ants-card {
    border-radius: var(--radius-lg);
    transition:
      box-shadow var(--duration-medium) var(--easing-default),
      background-color var(--duration-medium) var(--easing-default);
  }

  .ants-card-header {
    padding: 16px 16px 0;
  }

  .ants-card-body {
    width: 100%;
  }

  .ants-card-body--padded {
    padding: 16px;
  }

  .ants-card-footer {
    padding: 0 16px 16px;
  }

  /* ---- Default (card with border + subtle shadow) ---- */
  .ants-card--default {
    background-color: var(--bg-layer-alt);
    border: 1px solid var(--stroke-card);
    box-shadow: var(--shadow-2);
  }

  /* ---- Filled (solid background, no border) ---- */
  .ants-card--filled {
    background-color: var(--bg-layer);
    box-shadow: var(--shadow-2);
  }

  /* ---- Outlined (border only, transparent bg) ---- */
  .ants-card--outlined {
    background-color: transparent;
    border: 1px solid var(--stroke-default);
  }

  /* ---- Acrylic (translucent background) ---- */
  .ants-card--acrylic {
    background-color: var(--bg-acrylic);
    border: 1px solid var(--stroke-subtle);
    box-shadow: var(--shadow-4);
    backdrop-filter: blur(20px) saturate(125%);
    -webkit-backdrop-filter: blur(20px) saturate(125%);
  }
</style>
