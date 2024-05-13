<svelte:head>
	<title>About</title>
	<meta name="description" content="About this app" />
</svelte:head>

<script>
	import { page } from '$app/stores';
	import { goto } from '$app/navigation';
	import { NewProject2} from "$lib/wailsjs/go/main/App";
	import { Label, Input, Button, InputAddon, ButtonGroup, Checkbox,Avatar,Dropdown, DropdownItem, Search } from 'flowbite-svelte';
	function load1(){
		NewProject2();
	}

</script>

<div class="text-column">
	<h1>Projekt</h1>


		<Button color="dark" on:click={load1}>Laden</Button>


</div>