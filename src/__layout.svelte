<script>
  import { onMount } from 'svelte';

  onMount(() => {
    if (window.matchMedia && window.matchMedia('(prefers-color-scheme: light)').matches) {
      document.documentElement.classList.add('light-mode');
    }
  });
</script>

<slot />