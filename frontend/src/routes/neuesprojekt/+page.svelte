<svelte:head>
	<title>About</title>
	<meta name="description" content="About this app" />
</svelte:head>
<script>
	import { ListTrees, NewProject} from "$lib/wailsjs/go/main/App";
	import { Button, Modal, Label, Input, Checkbox, Dropdown, DropdownItem, Select } from 'flowbite-svelte';
	import { ChevronDownOutline } from 'flowbite-svelte-icons';
  	let formModal = false;
	let kunde = "";
	let projektnummer = "";
	let projektbeschreibung = "";
	let jahr = "";
	let jahrselected = [
    { value: "2024", name: "2024" }
  	];
	let kundeselected = [
    { value: "KROENERT", name: "KROENERT" },
    { value: "Ludlum", name: "Ludlum" },
    { value: "MFH", name: "MF Hamburg" }
  	];

	  function createProjekt(){
		NewProject(jahr,kunde, projektnummer + "_" + projektbeschreibung);
	}
</script>
<Button color="dark" on:click={() => (formModal = true)}>Neues Projekt</Button>
<Modal bind:open={formModal} size="xs" autoclose={false} class="w-full">
	<form class="flex flex-col space-y-6" action="#">
	  <h3 class="mb-4 text-xl font-medium text-gray-900 dark:text-white">Projektdaten ausfüllen</h3>
	  <Label>
		Kunde
		<Select class="mt-2" items={kundeselected} bind:value={kunde} required/>
	  </Label>
	  <Label class="space-y-2">
		<span>Projektnummer</span>
		<Input type="text" name="projektnummer" placeholder="projektnummer" bind:value={projektnummer} required />
	  </Label>
	  <Label class="space-y-2">
		<span>Beschreibung</span>
		<Input type="text" name="projektbeschreibung" placeholder="projektbeschreibung" bind:value={projektbeschreibung} required />
	  </Label>
	  <Label>
		Jahr
		<Select class="mt-2" items={jahrselected} bind:value={jahr} required/>
	  </Label>
	  <Button color="dark" type="submit" class="w-full1" on:click={createProjekt}>Projekt erstellen</Button>
	</form>
  </Modal>
