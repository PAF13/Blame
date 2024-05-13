<script>
	import "../app.pcss";
    import Header from './Header.svelte';
    import './styles.css';
	import { SpeedDial, SpeedDialButton } from 'flowbite-svelte';
	import { ListTrees, NewProject} from "$lib/wailsjs/go/main/App";
	import { Button, Modal, Label, Input, Checkbox, Dropdown, DropdownItem, Select } from 'flowbite-svelte';
	import { ChevronDownOutline } from 'flowbite-svelte-icons';

	import { Drawer, CloseButton, Sidebar, SidebarBrand, SidebarCta, SidebarDropdownItem, SidebarDropdownWrapper, SidebarGroup, SidebarItem, SidebarWrapper } from 'flowbite-svelte';
  import { ChartPieSolid, GridSolid, MailBoxSolid, UsersSolid, ArrowRightToBracketOutline,} from 'flowbite-svelte-icons';
  import { sineIn } from 'svelte/easing';
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
	
	let hidden2 = true;
  let spanClass = 'flex-1 ms-3 whitespace-nowrap';
  let transitionParams = {
    x: -320,
    duration: 200,
    easing: sineIn
  };
</script>

<div class="app">
	<Header></Header>

	<main>
		<slot></slot>
	</main>

	<SpeedDial bind:open defaultClass="absolute end-6 bottom-6" color="red">
		<SpeedDialButton name="Neues Projekt"  on:click={close} on:click={() => (formModal = true)}>Projekt</SpeedDialButton>
		<SpeedDialButton name="Print" on:click={() => (hidden2 = false)}>Show navigation</SpeedDialButton>
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

  <Drawer transitionType="fly" {transitionParams} bind:hidden={hidden2} id="sidebar2">
	<div class="flex items-center">
	  <h5 id="drawer-navigation-label-3" class="text-base font-semibold text-gray-500 uppercase dark:text-gray-400">Menu</h5>
	  <CloseButton on:click={() => (hidden2 = true)} class="mb-4 dark:text-white" />
	</div>
	<Sidebar>
	  <SidebarWrapper divClass="overflow-y-auto py-4 px-3 rounded dark:bg-gray-800">
		<SidebarGroup>
		  <SidebarItem label="Dashboard">
			<svelte:fragment slot="icon">
			  <ChartPieSolid class="w-5 h-5 text-gray-500 transition duration-75 dark:text-gray-400 group-hover:text-gray-900 dark:group-hover:text-white" />
			</svelte:fragment>
		  </SidebarItem>
		  <SidebarItem label="Kanban" {spanClass}>
			<svelte:fragment slot="icon">
			  <GridSolid class="w-5 h-5 text-gray-500 transition duration-75 dark:text-gray-400 group-hover:text-gray-900 dark:group-hover:text-white" />
			</svelte:fragment>
			<svelte:fragment slot="subtext">
			  <span class="inline-flex justify-center items-center px-2 ms-3 text-sm font-medium text-gray-800 bg-gray-200 rounded-full dark:bg-gray-700 dark:text-gray-300"> Pro </span>
			</svelte:fragment>
		  </SidebarItem>
		  <SidebarItem label="Inbox" {spanClass}>
			<svelte:fragment slot="icon">
			  <MailBoxSolid class="w-5 h-5 text-gray-500 transition duration-75 dark:text-gray-400 group-hover:text-gray-900 dark:group-hover:text-white" />
			</svelte:fragment>
			<svelte:fragment slot="subtext">
			  <span class="inline-flex justify-center items-center p-3 ms-3 w-3 h-3 text-sm font-medium text-primary-600 bg-primary-200 rounded-full dark:bg-primary-900 dark:text-primary-200"> 3 </span>
			</svelte:fragment>
		  </SidebarItem>
		  <SidebarItem label="Users">
			<svelte:fragment slot="icon">
			  <UsersSolid class="w-5 h-5 text-gray-500 transition duration-75 dark:text-gray-400 group-hover:text-gray-900 dark:group-hover:text-white" />
			</svelte:fragment>
		  </SidebarItem>
		  <SidebarItem label="Sign In">
			<svelte:fragment slot="icon">
			  <ArrowRightToBracketOutline class="w-5 h-5 text-gray-500 transition duration-75 dark:text-gray-400 group-hover:text-gray-900 dark:group-hover:text-white" />
			</svelte:fragment>
		  </SidebarItem>
		</SidebarGroup>
	  </SidebarWrapper>
	</Sidebar>
  </Drawer>

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
		width:  100vw;
		max-width: 100vw;
		margin: 0 auto;
		box-sizing: border-box;
	}
</style>
