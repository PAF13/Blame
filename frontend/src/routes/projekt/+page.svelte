<svelte:head>
	<title>About</title>
	<meta name="description" content="About this app" />
</svelte:head>

<script>
	import { page } from '$app/stores';
	import { goto } from '$app/navigation';
	import { NewProject2,CurrentProject,AddProdukt} from "$lib/wailsjs/go/main/App";
	import { Label, Input, Button, InputAddon, ButtonGroup, Checkbox,Avatar,Dropdown, DropdownItem, Search } from 'flowbite-svelte';
	function load1(){
		NewProject2();
	}
	function load2(){
		CurrentProject();
	}
	function load3(){
		AddProdukt();
	}
</script>

<div class="text-column">
	<h1>Projekt</h1>


		<Button color="dark" on:click={load1}>Laden</Button>
		<Button color="dark" on:click={load2}>Current Project</Button>
		<Button color="dark" on:click={load3}>add Produkt</Button>

</div>