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

<div class="winui-input-group">
  {#if label}
    <label for={id} class="winui-label" class:winui-label--disabled={disabled}>
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
    class="winui-input winui-input--{size}"
    class:winui-input--error={!!error}
    class:winui-input--disabled={disabled}
  />
  {#if error}
    <span class="winui-error">{error}</span>
  {/if}
</div>

<style>
  .winui-input-group {
    display: flex;
    flex-direction: column;
    gap: 4px;
  }

  .winui-label {
    font-size: var(--text-body);
    line-height: var(--text-body-lh);
    font-weight: var(--weight-semibold);
    color: var(--text-primary);
  }

  .winui-label--disabled {
    color: var(--text-disabled);
  }

  .winui-input {
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

  .winui-input::placeholder {
    color: var(--text-tertiary);
  }

  .winui-input:hover:not(:disabled) {
    border-color: var(--stroke-strong);
  }

  .winui-input:focus {
    border-color: var(--accent-default);
    box-shadow: 0 0 0 1px var(--accent-default);
  }

  .winui-input:disabled {
    opacity: 0.38;
    cursor: not-allowed;
    background-color: var(--stroke-subtle);
  }

  /* ---- Sizes ---- */
  .winui-input--sm {
    padding: 4px 10px;
    font-size: var(--text-body-small);
    line-height: var(--text-body-small-lh);
    height: 28px;
  }

  .winui-input--md {
    padding: 5px 12px;
    font-size: var(--text-body);
    line-height: var(--text-body-lh);
    height: 32px;
  }

  .winui-input--lg {
    padding: 7px 14px;
    font-size: var(--text-body);
    line-height: var(--text-body-lh);
    height: 40px;
  }

  /* ---- Error State ---- */
  .winui-input--error {
    border-color: var(--status-critical);
  }

  .winui-input--error:focus {
    border-color: var(--status-critical);
    box-shadow: 0 0 0 1px var(--status-critical);
  }

  .winui-error {
    font-size: var(--text-caption);
    line-height: var(--text-caption-lh);
    color: var(--status-critical-text);
  }
</style>
