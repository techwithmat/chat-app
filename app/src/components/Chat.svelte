<script lang="ts">
	import { username } from '../stores'
	import ChatHeader from './ChatHeader.svelte'
	import ChatFooter from './ChatFooter.svelte'
	import ChatBody from './ChatBody.svelte'

	let socket = new WebSocket('ws://localhost:3000/ws')

	socket.onopen = () => {
		socket.send(JSON.stringify({ action: 'connect', username: $username }))
	}
</script>

<div class="h-full grid grid-rows-[60px,1fr,60px] gap-4">
	<ChatHeader connection={socket} />
	<ChatBody connection={socket} />
	<ChatFooter connection={socket} />
</div>
