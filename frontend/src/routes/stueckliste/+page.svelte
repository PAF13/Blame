<svelte:head>
	<title>About</title>
	<meta name="description" content="About this app" />
</svelte:head>
<script lang="ts">
	import { Label, Select,MultiSelect, Input,ButtonGroup,Button} from 'flowbite-svelte';
	import { ChevronDownOutline, UserRemoveSolid, SearchOutline} from 'flowbite-svelte-icons';
	import { OpenMultipleFilesDialog,LoadStueckliste,ReturnOrte} from "$lib/wailsjs/go/main/App";


  let searchTerm = ''
  const kunde = [
	{ value: 'KNT', name: 'Kroenert' },
    { value: 'TOPIX', name: 'Topx' },
    { value: 'SITECA', name: 'Siteca' }
	];
	let selectedKunde: string;

	let orte = [
    { value: 'us', name: 'United States' },
    { value: 'ca', name: 'Canada' },
    { value: 'fr', name: 'France' }
  	];
	let selectedOrte: string[];


  $: filteredItems = kunde.filter((person) => person.name.toLowerCase().indexOf(searchTerm?.toLowerCase()) !== -1);

  	let pfad = "";
	let pfad1Last = "";
	let pfad1Switch = "";

	let pfad1Loaded = Boolean(false);

	let pfad2 = "";
	let pfad2Last = "";
	let pfad2Switch = "";
	let pfad2Loaded = Boolean(false);


	let BMK = new Array();
	let BMK2 = new Array();
	
	let pfaden: string[];
	function load1(){
		OpenMultipleFilesDialog().then((result) => (pfaden = result));
	}

	
	function dialog2(){
		LoadStueckliste(pfaden, selectedKunde, "stueckliste");
	}





	function test(){
		ReturnOrte().then((result) => (BMK = result));
	}






</script>


<h1>Stückliste</h1>
<label for="">{pfaden}</label>
<ButtonGroup>
	<Button color="dark" on:click={load1}>Datei...</Button>
	<Select color="dark" items={kunde} bind:value={selectedKunde} />
	<Button color="dark" on:click={dialog2}>Profile</Button>
</ButtonGroup>

<MultiSelect items={orte} bind:value={selectedOrte} />