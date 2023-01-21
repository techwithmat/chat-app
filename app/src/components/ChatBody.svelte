<script lang="ts">
	export let connection: WebSocket
	import { messages } from '../stores'
	import Message from './Message.svelte'

	connection.onmessage = (e) => {
		let data = JSON.parse(e.data)

		messages.update((msg) => [...msg, data])
	}
</script>

<div
	class="flex flex-col gap-2 overflow-auto scrollbar-thin scrollbar-thumb-zinc-500 hover:scrollbar-thumb-zinc-400"
>
	{#if $messages}
		{#each $messages as message}
			<Message message={message.message} user={message.username} />
		{/each}
	{/if}
</div>
