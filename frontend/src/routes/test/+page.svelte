<svelte:head>
	<title>About</title>
	<meta name="description" content="About this app" />
</svelte:head>
<script>
	import { Dialog, VerbindungCorrection } from "$lib/wailsjs/go/main/App";
	import { Label, Listgroup, ListgroupItem, ButtonGroup, Button, Input} from 'flowbite-svelte';
	import {Dropdown, DropdownItem, Checkbox, Search } from 'flowbite-svelte';
  import { ChevronDownOutline, UserRemoveSolid } from 'flowbite-svelte-icons';

	let num = 0;


	function sliceAppending(){
		VerbindungCorrection();
	}
	const handleClick = (e) => {
    e.preventDefault();
    alert('Clicked on: ' + e.target);
  };
  </script>
  
  <div class="pt-8">
	<Label for="input-addon" class="mb-2">Alte Stückliste</Label>
	<ButtonGroup class="w-full">
		<Button color="dark" on:click={sliceAppending}>importieren</Button>
		<Input id="input-addon" type="number" value={num} placeholder="Dateipfad" />
	</ButtonGroup>
	</div>

	<Button color="dark" >Dropdown search<ChevronDownOutline class="w-3 h-3 ms-2 text-white dark:text-white" /></Button>
<Dropdown class="overflow-y-auto px-3 pb-3 text-sm h-44">
  <div slot="header" class="p-3">
    <Search size="md" />
  </div>
  <li class="rounded p-2 hover:bg-gray-100 dark:hover:bg-gray-600">
    <Checkbox>Robert Gouth</Checkbox>
  </li>
  <li class="rounded p-2 hover:bg-gray-100 dark:hover:bg-gray-600">
    <Checkbox>Jese Leos</Checkbox>
  </li>
  <li class="rounded p-2 hover:bg-gray-100 dark:hover:bg-gray-600">
    <Checkbox checked>Bonnie Green</Checkbox>
  </li>
  <li class="rounded p-2 hover:bg-gray-100 dark:hover:bg-gray-600">
    <Checkbox>Jese Leos</Checkbox>
  </li>
  <li class="rounded p-2 hover:bg-gray-100 dark:hover:bg-gray-600">
    <Checkbox>Robert Gouth</Checkbox>
  </li>
  <li class="rounded p-2 hover:bg-gray-100 dark:hover:bg-gray-600">
    <Checkbox>Bonnie Green</Checkbox>
  </li>
  <a slot="footer" href="/" class="flex items-center p-3 -mb-1 text-sm font-medium text-red-600 bg-gray-50 hover:bg-gray-100 dark:bg-gray-700 dark:hover:bg-gray-600 dark:text-red-500 hover:underline">
    <UserRemoveSolid class="w-4 h-4 me-2 text-primary-700 dark:text-primary-700" />Delete user
  </a>
</Dropdown>


<Button color="dark" >Dropdown button<ChevronDownOutline class="w-3 h-3 ms-2 text-white dark:text-white" /></Button>
<Dropdown>
  <DropdownItem href="/verbindungsliste" on:click={handleClick}>Rendered as link</DropdownItem>
  <DropdownItem on:click={handleClick}>Rendered as button</DropdownItem>
</Dropdown>