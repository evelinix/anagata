<script lang="ts">
  type InputType = 'text' | 'password' | 'email' | 'number' | 'search' | 'url';
  type Size = 'sm' | 'md' | 'lg';

  let {
    type = 'text',
    size = 'md',
    label = '',
    placeholder = '',
    value = $bindable(''),
    error = '',
    disabled = false,
    readonly = false,
    name = '',
    id = '',
  }: {
    type?: InputType;
    size?: Size;
    label?: string;
    placeholder?: string;
    value?: string;
    error?: string;
    disabled?: boolean;
    readonly?: boolean;
    name?: string;
    id?: string;
  } = $props();
</script>

<div class="ants-input-group">
  {#if label}
    <label for={id} class="ants-label" class:ants-label--disabled={disabled}>
      {label}
    </label>
  {/if}
  <input
    {type}
    {name}
    {id}
    {placeholder}
    {disabled}
    {readonly}
    bind:value
    class="ants-input ants-input--{size}"
    class:ants-input--error={!!error}
    class:ants-input--disabled={disabled}
  />
  {#if error}
    <span class="ants-error">{error}</span>
  {/if}
</div>

<style>
  .ants-input-group {
    display: flex;
    flex-direction: column;
    gap: 4px;
  }

  .ants-label {
    font-size: var(--text-body);
    line-height: var(--text-body-lh);
    font-weight: var(--weight-semibold);
    color: var(--text-primary);
  }

  .ants-label--disabled {
    color: var(--text-disabled);
  }

  .ants-input {
    font-family: var(--font-family);
    color: var(--text-primary);
    background-color: var(--bg-layer-alt);
    border: 1px solid var(--stroke-default);
    border-radius: var(--radius-md);
    outline: none;
    transition:
      border-color var(--duration-short) var(--easing-default),
      background-color var(--duration-short) var(--easing-default),
      box-shadow var(--duration-short) var(--easing-default);
  }

  .ants-input::placeholder {
    color: var(--text-tertiary);
  }

  .ants-input:hover:not(:disabled) {
    border-color: var(--stroke-strong);
  }

  .ants-input:focus {
    border-color: var(--accent-default);
    box-shadow: 0 0 0 1px var(--accent-default);
  }

  .ants-input:disabled {
    opacity: 0.38;
    cursor: not-allowed;
    background-color: var(--stroke-subtle);
  }

  /* ---- Sizes ---- */
  .ants-input--sm {
    padding: 4px 10px;
    font-size: var(--text-body-small);
    line-height: var(--text-body-small-lh);
    height: 28px;
  }

  .ants-input--md {
    padding: 5px 12px;
    font-size: var(--text-body);
    line-height: var(--text-body-lh);
    height: 32px;
  }

  .ants-input--lg {
    padding: 7px 14px;
    font-size: var(--text-body);
    line-height: var(--text-body-lh);
    height: 40px;
  }

  /* ---- Error State ---- */
  .ants-input--error {
    border-color: var(--status-critical);
  }

  .ants-input--error:focus {
    border-color: var(--status-critical);
    box-shadow: 0 0 0 1px var(--status-critical);
  }

  .ants-error {
    font-size: var(--text-caption);
    line-height: var(--text-caption-lh);
    color: var(--status-critical-text);
  }
</style>
