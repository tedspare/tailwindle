<script>
  import { enhance } from "$lib/form";
  import { scale } from "svelte/transition";
  import { flip } from "svelte/animate";

  export let todos;
</script>

<svelte:head>
  <title>Todos</title>
</svelte:head>

<div class="w-full text-center">
  <h1>Todos</h1>
  <form
    action="/todos"
    method="post"
    use:enhance={{
      result: async ({ form }) => {
        form.reset();
      },
    }}
  >
    <input
      name="text"
      aria-label="Add todo"
      placeholder="Add a todo"
      class="p-3 rounded-lg text-gray-600"
    />
  </form>
  {#each todos as todo (todo.uid)}
    <div
      class:opacity-40={todo.done}
      transition:scale|local={{ start: 0.7 }}
      animate:flip={{ duration: 200 }}
      class="flex items-center justify-center pt-6"
    >
      <form
        action="/todos?_method=PATCH"
        method="post"
        use:enhance={{
          pending: ({ data }) => {
            todo.done = !!data.get("done");
          },
        }}
      >
        <input type="hidden" name="uid" value={todo.uid} />
        <input type="hidden" name="done" value={todo.done ? "" : "true"} />
        <button
          aria-label="Mark todo as {todo.done ? 'not done' : 'done'}"
          class="text-2xl"
        >
          ✅
        </button>
      </form>
      <form
        action="/todos?_method=PATCH"
        method="post"
        class="text-black p-3 rounded-lg bg-gray-100"
        use:enhance
      >
        <input type="hidden" name="uid" value={todo.uid} />
        <input
          aria-label="Edit todo"
          type="text"
          name="text"
          class="h-10 bg-gray-100"
          value={todo.text}
        />
        <button aria-label="Save todo" class="bg-blue-300">Save</button>
      </form>
      <form
        action="/todos?_method=DELETE"
        method="post"
        use:enhance={{
          pending: () => (todo.pending_delete = true),
        }}
      >
        <input type="hidden" name="uid" value={todo.uid} />
        <button
          aria-label="Delete todo"
          class="bg-red-600 ml-2"
          disabled={todo.pending_delete}
        >
          Delete
        </button>
      </form>
    </div>
  {/each}
  {#if todos.every((todo) => todo.done)}
    <div class="pt-4">All done!</div>
  {/if}
</div>
