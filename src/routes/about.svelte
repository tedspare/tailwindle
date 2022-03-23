<script context="module">
  import { browser, dev } from "$app/env";

  // we don't need any JS on this page, though we'll load
  // it in dev so that we get hot module replacement...
  export const hydrate = dev;

  // ...but if the client-side router is already loaded
  // (i.e. we came here from elsewhere in the app), use it
  export const router = browser;

  // since there's no dynamic data here, we can prerender
  // it so that it gets served as a static asset in prod
  export const prerender = true;
</script>

<svelte:head>
  <title>About</title>
</svelte:head>

<div class="w-full max-w-3xl mx-auto text-center">
  <h1>Wordle for Tailwind CSS</h1>
  <p>
    That's the idea. You get six attempts to guess five <a
      href="https://tailwindcss.com/">Tailwind</a
    >
    classes applied to a
    <a href="https://www.w3schools.com/tags/tag_div.asp">div</a>.
  </p>
  <p>More to come.</p>
</div>
