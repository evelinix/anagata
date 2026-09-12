<script>
  let { children, fallback } = $props();
  let error = $state(null);
</script>

{#if error}
  {#if fallback}
    {@render fallback(error)}
  {:else}
    <div class="error-boundary">
      <h3>Something went wrong</h3>
      <p>{error.message || 'An unexpected error occurred'}</p>
      <button onclick={() => (error = null)}>Try again</button>
    </div>
  {/if}
{:else}
  {@render children()}
{/if}

<style>
  .error-boundary {
    padding: 1rem;
    border: 1px solid #ef4444;
    border-radius: 0.5rem;
    background: #fef2f2;
    color: #991b1b;
  }
  button {
    margin-top: 0.5rem;
    padding: 0.5rem 1rem;
    background: #ef4444;
    color: white;
    border: none;
    border-radius: 0.25rem;
    cursor: pointer;
  }
</style>
