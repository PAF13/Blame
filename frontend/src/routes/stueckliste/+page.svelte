<svelte:head>
	<title>About</title>
	<meta name="description" content="About this app" />
</svelte:head>
<script>
	import { Label, Input, Button, InputAddon, ButtonGroup, Checkbox,Avatar,Dropdown, DropdownItem, Search } from 'flowbite-svelte';
	import { ChevronDownOutline, UserRemoveSolid } from 'flowbite-svelte-icons';
	import { OpenFileDialog,LoadStueckliste} from "$lib/wailsjs/go/main/App";
  let searchTerm = ''
  const people = [{ name: 'All', checked: false }];
  const people2 = new Array();


  $: filteredItems = people.filter((person) => person.name.toLowerCase().indexOf(searchTerm?.toLowerCase()) !== -1);

  let list1 = "";
	let list1Switch = "";
	let list1Last = "replace";
	let list2 = "";
	let list2Switch = "";
	let list2Last = "replace";

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
		if(list1 != list1Last){
			list1 = "loading..."
			list1Last = "loading..."
			OpenFileDialog().then((result) => (list1 = result));
		}
	}

	$: {
		if (BMK != BMK2) {
			ii2 = BMK.length;
		for (let i = 0; i < ii2; i++) {
			people[0]["name"] = "All";
			people[0]["checked"] = true;
			people.push({ name: BMK[i], checked: false });
			
			}
			BMK2 = BMK;
		}
	}

	function dialog2(){
		LoadStueckliste(list1);

	}

	function reset(){
		pfad1 = "";
		pfad2 = "";
		pfad1Loaded = false;
		pfad2Loaded = false;
	}



function switchList(){
	pfad1Switch = pfad1;
	pfad2Switch = pfad2;
	pfad2 = pfad1Switch;
	pfad1 = pfad2Switch;
}

</script>


<h1>Stückliste</h1>

<div class="pt-8">
<ButtonGroup class="w-full">
	<Button color="dark" on:click={dialog1}>importieren</Button>
	<Button color="dark" on:click={dialog2}>Laden</Button>
	<Input 		
		id="input-addon" 
		type="text" 
		value={list1} 
		placeholder="Alte Stückliste" 
		color='green'/>
</ButtonGroup>
</div>

	<Button color="dark" >Dropdown search<ChevronDownOutline class="w-6 h-6 ms-2 text-white dark:text-white" /></Button>
	<Dropdown class="overflow-y-auto px-3 pb-3 text-sm h-44">
		<div slot="header" class="p-3">
		  <Search size="md" bind:value={searchTerm}/>
		</div>
		
		{#each filteredItems as person (person.name)}
		
		  <DropdownItem class="flex items-center text-base font-semibold gap-2">
			
			<Checkbox bind:checked={person.checked}>{person.name}</Checkbox>
			
			</DropdownItem>
		{/each}
		<a slot="footer" href="/" class="flex items-center px-3 py-2 -mb-1 text-sm font-medium text-primary-600 bg-gray-50 hover:bg-gray-100 dark:bg-gray-700 dark:hover:bg-gray-600 dark:text-primary-500 hover:underline">
			<UserRemoveSolid class="w-4 h-4 me-2 text-primary-700 dark:text-primary-700" />Delete user
		</a>
	  </Dropdown>

