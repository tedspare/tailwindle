<script context="module">
  export const prerender = true;
</script>

<script>
  import Counter from "$lib/Counter.svelte";
</script>

<svelte:head>
  <title>Tailwindle</title>
</svelte:head>

<section class="flex flex-col justify-center items-center">
  <h1 class="">Welcome to Tailwindle</h1>
  <h2>
    Try editing <strong>src/routes/index.svelte</strong>
  </h2>
  <Counter />
</section>
