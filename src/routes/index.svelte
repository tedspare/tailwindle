<script context="module">
	export const prerender = true;
</script>

<script>
	import Counter from '$lib/Counter.svelte';
</script>

<svelte:head>
	<title>Home</title>
</svelte:head>

<section class="flex flex-col justify-center items-center flex-1">
	<h1 class="text-blue-400 font-extrabold w-full">
		Welcome to Tailwindle
	</h1>

	<h2>
		Try editing <strong>src/routes/index.svelte</strong>
	</h2>

	<Counter />
</section>
