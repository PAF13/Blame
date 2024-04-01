<script>
	import "../app.pcss";
    import Header from './Header.svelte';
    import './styles.css';
	import { SpeedDial, SpeedDialButton } from 'flowbite-svelte';
	import { ListTrees, NewProject} from "$lib/wailsjs/go/main/App";
	import { Button, Modal, Label, Input, Checkbox, Dropdown, DropdownItem, Select } from 'flowbite-svelte';
	import { ChevronDownOutline } from 'flowbite-svelte-icons';
  let open = false;

  const close = () => {
    open = false;
  };
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

<div class="app">
	<Header></Header>

	<main>
		<slot></slot>
	</main>

	<SpeedDial bind:open defaultClass="absolute end-6 bottom-6" color="red">
		<SpeedDialButton name="Neues Projekt"  on:click={close} on:click={() => (formModal = true)}>Projekt</SpeedDialButton>
		<SpeedDialButton name="Print" on:click={close}>Void</SpeedDialButton>
	  </SpeedDial>

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
</div>

<style>
	.app {
		display: flex;
		flex-direction: column;
		min-height: 100vh;
	}

	main {
		flex: 1;
		display: flex;
		flex-direction: column;
		padding: 1rem;
		width: 100%;
		max-width: 64rem;
		margin: 0 auto;
		box-sizing: border-box;
	}

	footer {
		display: flex;
		flex-direction: column;
		justify-content: center;
		align-items: center;
		padding: 12px;
	}

	footer a {
		font-weight: bold;
	}

	@media (min-width: 480px) {
		footer {
			padding: 12px 0;
		}
	}
</style>
