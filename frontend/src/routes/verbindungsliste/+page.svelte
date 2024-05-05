<svelte:head>
	<title>About</title>
	<meta name="description" content="About this app" />
</svelte:head>

<script>
	import { Dialog, ExcelChoice,  ImportStueckliste, StuecklisteSum, VerbindungRead ,StuecklisteSum2} from "$lib/wailsjs/go/main/App";
	import { Label, Input, Button, InputAddon, ButtonGroup, Checkbox,Avatar,Dropdown, DropdownItem, Search } from 'flowbite-svelte';
	import { ChevronDownOutline, UserRemoveSolid } from 'flowbite-svelte-icons';

	let pfad1 = "";
	let pfad1Last = "";
	let pfad1Switch = "";

	let pfad1Loaded = Boolean(false);

	let pfad2 = "";
	let pfad2Last = "";
	let pfad2Switch = "";
	let pfad2Loaded = Boolean(false);


	let pfad22 = Boolean(true);
	let defaultModal = Boolean(false);
	let ii = 0;
	let ii2 = 0;
	let BMK = new Array();
	let BMK2 = new Array();

	function dialog1(){
		Dialog().then((result) => (pfad1 = result));
		pfad2Loaded = true;
	}

	function load1(){
		VerbindungRead();
	}
</script>

<div class="text-column">
	<h1>Verbindungsliste</h1>


		<Button color="dark" on:click={load1}>Laden</Button>

</div>

