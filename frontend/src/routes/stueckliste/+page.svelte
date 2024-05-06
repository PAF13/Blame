<svelte:head>
	<title>About</title>
	<meta name="description" content="About this app" />
</svelte:head>
<script>
	import { Dialog} from "$lib/wailsjs/go/main/App";
	import { Label, Input, Button, InputAddon, ButtonGroup, Checkbox,Avatar,Dropdown, DropdownItem, Search } from 'flowbite-svelte';
	import { ChevronDownOutline, UserRemoveSolid } from 'flowbite-svelte-icons';
  let searchTerm = ''
  const people = [{ name: 'All', checked: false }];
  const people2 = new Array();


  $: filteredItems = people.filter((person) => person.name.toLowerCase().indexOf(searchTerm?.toLowerCase()) !== -1);

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
		let ii3 = 0;
		people.forEach(function (item, index) {

  			console.log(item, index);
			  if(item.checked == true){
				ii3++
				people2[ii3] = item.name;
				//people2.push(item.name);
			  }
		});

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
	<Button color="dark">Laden</Button>
	<Input 		
		id="input-addon" 
		type="text" 
		value={pfad1} 
		placeholder="Alte Stückliste" 
		color='green'/>
</ButtonGroup>
<ButtonGroup class="w-full">
	<Button color="dark" disabled={pfad22}>importieren</Button>
	<Button color="dark" disabled={pfad22}>Laden</Button>
	<Input id="input-addon" type="text" value={pfad2} placeholder="Neue Stückliste" color='green' disabled={pfad22}/>
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



<div class="pt-8">
	<ButtonGroup>
		<Button color="dark" disabled={defaultModal} on:click={reset}>Reset</Button>
		<Button color="dark" disabled={defaultModal} on:click={switchList}>Tauschen</Button>
	</ButtonGroup>
	</div>
<div class="pt-8">
	<ButtonGroup>
		<Button color="dark" disabled={defaultModal}>Differenz</Button>	
		<Button color="dark" disabled={defaultModal} on:click={dialog2}>Stueckliste Sum</Button>
		<Button color="dark" disabled={defaultModal}>verbindungsliste</Button>
	</ButtonGroup>
</div>
