<script lang="ts">
	import { username } from '../stores'
	export let user: string
	export let message: string
</script>

{#if user === 'System'}
	<p class="text-center text-lg font-medium">- {message} -</p>
{:else}
	<div
		class="flex flex-col rounded-lg bg-zinc-800 border border-zinc-700 p-2 whitespace-normal break-words"
	>
		<p class="mb-2">
			<span class="font-bold">{user === $username ? 'You' : user}</span> - Today
		</p>
		<p class="ml-1">{message}</p>
	</div>
{/if}
