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

<div class="winui-card winui-card--{variant}">
  {#if header}
    <div class="winui-card-header">
      {@render header()}
    </div>
  {/if}
  <div class="winui-card-body" class:winui-card-body--padded={padding}>
    {#if children}
      {@render children()}
    {/if}
  </div>
  {#if footer}
    <div class="winui-card-footer">
      {@render footer()}
    </div>
  {/if}
</div>

<style>
  .winui-card {
    border-radius: var(--radius-lg);
    transition:
      box-shadow var(--duration-medium) var(--easing-default),
      background-color var(--duration-medium) var(--easing-default);
  }

  .winui-card-header {
    padding: 16px 16px 0;
  }

  .winui-card-body {
    width: 100%;
  }

  .winui-card-body--padded {
    padding: 16px;
  }

  .winui-card-footer {
    padding: 0 16px 16px;
  }

  /* ---- Default (card with border + subtle shadow) ---- */
  .winui-card--default {
    background-color: var(--bg-layer-alt);
    border: 1px solid var(--stroke-card);
    box-shadow: var(--shadow-2);
  }

  /* ---- Filled (solid background, no border) ---- */
  .winui-card--filled {
    background-color: var(--bg-layer);
    box-shadow: var(--shadow-2);
  }

  /* ---- Outlined (border only, transparent bg) ---- */
  .winui-card--outlined {
    background-color: transparent;
    border: 1px solid var(--stroke-default);
  }

  /* ---- Acrylic (translucent background) ---- */
  .winui-card--acrylic {
    background-color: var(--bg-acrylic);
    border: 1px solid var(--stroke-subtle);
    box-shadow: var(--shadow-4);
    backdrop-filter: blur(20px) saturate(125%);
    -webkit-backdrop-filter: blur(20px) saturate(125%);
  }
</style>
