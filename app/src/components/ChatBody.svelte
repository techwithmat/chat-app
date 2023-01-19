<script lang="ts">
  import Message from './Message.svelte'
  // leo los mensaje y los renderizo
  const Messages = []
</script>

<div
  class="flex flex-col gap-2 overflow-auto scrollbar-thin scrollbar-thumb-zinc-500 hover:scrollbar-thumb-zinc-400"
>
  {#each Messages as item}
    <Message message={item.message} user={item.user} />
  {/each}
</div>
