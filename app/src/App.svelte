<script lang="ts">
  import { onDestroy } from 'svelte'
  import { username } from './stores'
  import Login from './components/Login.svelte'
  import Chat from './components/Chat.svelte'

  let user: string
  const unsubscribe = username.subscribe(value => {
    user = value
  })

  onDestroy(unsubscribe)
</script>

<main class="h-screen bg-zinc-900 text-white">
  <div
    class="h-screen max-w-sm mx-auto px-4 flex flex-col justify-center mb:border-4 mb:border-zinc-700 mb:rounded-2xl"
  >
    {#if !user}
      <Login />
    {:else}
      <Chat />
    {/if}
  </div>
</main>
