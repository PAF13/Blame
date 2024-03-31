<svelte:head>
	<title>About</title>
	<meta name="description" content="About this app" />
</svelte:head>
<script>
	import { ListTrees } from "$lib/wailsjs/go/main/App";
	let projektbeschreibung = '';
	let projektnummer = '';
	let jahr = '';
</script>

<div class="text-column">
	<h1>Neues Projekt</h1>
	<h3>Projektnummer</h3>
	<input bind:value={projektnummer} placeholder="enter your projektnummer" />
	<h3>Projektbeschreibung</h3>
	<input bind:value={projektbeschreibung} placeholder="enter your projektbeschreibung" />
	<h3>Jahr</h3>
	<input bind:value={jahr} placeholder="enter your jahr" />
	<button>List</button>
</div>
