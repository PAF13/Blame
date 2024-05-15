<svelte:head>
	<title>About</title>
	<meta name="description" content="About this app" />
</svelte:head>
<script lang="ts">
	import { Label, Select,MultiSelect, Input,ButtonGroup,Button,Dropdown,Search,Checkbox,P} from 'flowbite-svelte';
	import { ChevronDownOutline, UserRemoveSolid, SearchOutline,} from 'flowbite-svelte-icons';
	import { OpenMultipleFilesDialog,LoadStueckliste,ReturnOrte,ExportStueckliste} from "$lib/wailsjs/go/main/App";



  const kunde = [
	{ value: 'KNT', name: 'Kroenert' },
    { value: 'TOPIX', name: 'Topx' },
    { value: 'SITECA', name: 'Siteca' }
	];
	let selectedKunde: string;

	
	let orte = [
    { value: 'us', name: 'United States' },
  	];

	  let searchTerm = ''
	  
	  let people_TEMP: { name: string, checked: boolean }[] = [];

	  let people: { name: string, checked: boolean }[] = [];

  let selected= "";
  let countries = [
    { value: 'us', name: 'United States' },
    { value: 'ca', name: 'Canada' },
    { value: 'fr', name: 'France' }
  ];
 
  $: filteredItems = people.filter((person) => person.name.toLowerCase().indexOf(searchTerm?.toLowerCase()) !== -1);
	let selectedOrte = new Array();
	let test: boolean = false;
	let ort_TEMP: string[] =  [];
	let ort:string[] =  [];
	var length: number;
	var test2: boolean = false;
var num: number = 0;
var num2: number = 0;

	let BMK = new Array();

	
	let pfaden: string[] =  [];

	function load1(){
		OpenMultipleFilesDialog().then((result) => (pfaden = result));
		test = true;
	}

	
	function dialog2(){
		if (test && pfaden != null){
			LoadStueckliste(pfaden, selectedKunde, "stueckliste").then((result) => (ort_TEMP = result));
		}
		
		
		
	}
	$:{ if (ort != ort_TEMP){
		let check = 0;
		ort_TEMP.forEach(function (myans) {
			if (check > 0){
				people_TEMP.push({ name: myans, checked: true }); 
			}
			check++;				
		});
		ort = ort_TEMP,
		people = people_TEMP;
		test = false;
	}
	
}

	function dialog3(){
		selectedOrte.push(ort[0])
		for (let i = 0; i < people.length-1; i++) {
			if (people[i].checked){
				selectedOrte.push(people[i].name); 	
			}
		}	
		ExportStueckliste(selectedOrte, selectedKunde, "stueckliste");
	}







</script>


<h1>Stückliste</h1>


<label for="">{pfaden}</label>

<ButtonGroup>
	<Button color="dark" on:click={load1}>Datei...</Button>
	<Select color="dark" items={kunde} bind:value={selectedKunde} />
	<Button color="dark"  on:click={dialog2}>Dropdown search<ChevronDownOutline class="w-6 h-6 ms-2 text-white dark:text-white" /></Button>
<Dropdown class="overflow-y-auto px-3 pb-3 text-sm h-44">
  <div slot="header" class="p-3">
    <Search size="md" bind:value={searchTerm}/>
  </div>
  {#key test}
  {#each filteredItems as person (person.name)}
    <li class="rounded p-2 hover:bg-gray-100 dark:hover:bg-gray-600">
      <Checkbox bind:checked={person.checked}>{person.name}</Checkbox>
    </li>
  {/each}
  {/key}
</Dropdown>
<Button color="dark" on:click={dialog3}>Export</Button>
</ButtonGroup>
<h1>Orten</h1>
{#key people}
  {#each filteredItems as person (person.name)}
		<P size="xl" linethrough opacity={25} color="text-blue-600 dark:text-blue-500">{person.name}</P>
  {/each}
  {/key}

