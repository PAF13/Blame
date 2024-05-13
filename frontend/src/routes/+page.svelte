<svelte:head>
	<title>Home</title>
	<meta name="description" content="Svelte demo app" />
</svelte:head>
<script>
	import { AccordionItem, Accordion, Button,Tabs, TabItem,Table, TableBody, TableBodyCell, TableBodyRow, TableHead, TableHeadCell, TableSearch, Checkbox    } from 'flowbite-svelte';
 	const items2 = Array(3);

  	const open_all = () => items.forEach((_, i) => (items2[i] = true));
  	const close_all = () => items.forEach((_, i) => (items2[i] = false));

	  let searchTerm = '';
  let items = [
    { Typ: "Schaltschrank", BMK: '+EC0', Bearbeiter_SITECA: 'MMA', Bearbeiter_Kunde: "Martin Nohr" , Liefertermin: "27/05/24", Prio: "", konstruktionsvorgaben: "", eplan: "", pruefung: "", freigabe: "", NC: "", material: "", eview : "", verbindungsliste: "", klemmenaufreihplan: "", stromlaufplan: ""},
	{ Typ: "Schaltschrank", BMK: '+EC1.0', Bearbeiter_SITECA: 'MMA', Bearbeiter_Kunde: "Martin Nohr" , Liefertermin: "27/05/24", Prio: ""},
	{ Typ: "Schaltschrank", BMK: '+EC1', Bearbeiter_SITECA: 'MMA', Bearbeiter_Kunde: "Martin Nohr" , Liefertermin: "27/05/24", Prio: ""},
	{ Typ: "Schaltschrank", BMK: '+EC2', Bearbeiter_SITECA: 'MMA', Bearbeiter_Kunde: "Martin Nohr" , Liefertermin: "27/05/24", Prio: ""},
	{ Typ: "Schaltschrank", BMK: '+EC3', Bearbeiter_SITECA: 'MMA', Bearbeiter_Kunde: "Martin Nohr" , Liefertermin: "27/05/24", Prio: ""},
	{ Typ: "Schaltschrank", BMK: '+EC4', Bearbeiter_SITECA: 'MMA', Bearbeiter_Kunde: "Martin Nohr" , Liefertermin: "27/05/24", Prio: ""},
	{ Typ: "Schaltschrank", BMK: '+EC5', Bearbeiter_SITECA: 'MMA', Bearbeiter_Kunde: "Martin Nohr" , Liefertermin: "27/05/24", Prio: ""},
  ];
  $: filteredItems = items.filter((item) => item.BMK.toLowerCase().indexOf(searchTerm.toLowerCase()) !== -1);
  let w=3;
</script>


<section>
	

	<h1>Home</h1>
</section>
<TableSearch striped={true} placeholder="Search by maker name" hoverable={true} bind:inputValue={searchTerm}>
  <Tabs>
	<TabItem open title="Konstruktion">

		<TableHead>
			<TableHeadCell class="!p-4">
				<Checkbox />
			</TableHeadCell>
		  <TableHeadCell>Typ</TableHeadCell>
		  <TableHeadCell>BMK/Einbauort</TableHeadCell>
		  <TableHeadCell>Bearbeiter SITECA</TableHeadCell>
		  <TableHeadCell>Bearbeiter Kunde</TableHeadCell>
		  <TableHeadCell>Liefertermin</TableHeadCell>
		  <TableHeadCell>Hinweis / Prio</TableHeadCell>
		  <TableHeadCell>Konstruktionsvorgaben prüfen</TableHeadCell>
		  <TableHeadCell>Prüfung 4 Augen</TableHeadCell>
		  <TableHeadCell>FREIGABE Kunde</TableHeadCell>
		  <TableHeadCell>Aufbau / NC Daten</TableHeadCell>
		  <TableHeadCell>Beschriftung BMK</TableHeadCell>
		  <TableHeadCell>Artikelsummenstückliste / Material</TableHeadCell>
		  <TableHeadCell>Eview hochgeladen</TableHeadCell>
		  <TableHeadCell>Verbindungsliste</TableHeadCell>
		  <TableHeadCell>Klemmenaufreihplan/ Klemmennummer</TableHeadCell>
		  <TableHeadCell>Stromlaufplan</TableHeadCell>
		</TableHead>
		<TableBody tableBodyClass="divide-y">
		  {#each filteredItems as item}
			<TableBodyRow>
				<TableBodyCell class="!p-4">
					<Checkbox />
				</TableBodyCell>
	
			  <TableBodyCell>{item.Typ}</TableBodyCell>
			  <TableBodyCell>{item.BMK}</TableBodyCell>
			  <TableBodyCell>{item.Bearbeiter_SITECA}</TableBodyCell>
			  <TableBodyCell>{item.Bearbeiter_Kunde}</TableBodyCell>
			  <TableBodyCell>{item.Liefertermin}</TableBodyCell>
			  <TableBodyCell>{item.Prio}</TableBodyCell>
			</TableBodyRow>
			
		  {/each}
		</TableBody>
		<tfoot>
			<tr class="font-semibold text-gray-900 dark:text-white">
			  <th scope="row" class="py-3 px-6 text-base">Total</th>
			  <td class="py-3 px-6">3</td>
			  <td class="py-3 px-6">21,000</td>
			</tr>
		  </tfoot>
	</TabItem>
	<TabItem title="Einkauf">
		<TableHead>
			<TableHeadCell class="!p-4">
				<Checkbox />
			</TableHeadCell>
		  <TableHeadCell>Typ</TableHeadCell>
		  <TableHeadCell>BMK/Einbauort</TableHeadCell>
		  <TableHeadCell>Bearbeiter SITECA</TableHeadCell>
		  <TableHeadCell>Bearbeiter Kunde</TableHeadCell>
		  <TableHeadCell>Liefertermin</TableHeadCell>
		  <TableHeadCell>Hinweis / Prio</TableHeadCell>
		  
		  <TableHeadCell>Summenstückliste bearbeiten</TableHeadCell>
		  <TableHeadCell>Summenstückliste einlesen Topix</TableHeadCell>
		  <TableHeadCell>Bestellungen auslösen Topix</TableHeadCell>
		  <TableHeadCell>Material termingerecht</TableHeadCell>
		  <TableHeadCell>Änderungen an Konstruktion übermittelt</TableHeadCell>
		</TableHead>
		<TableBody tableBodyClass="divide-y">
		  {#each filteredItems as item}
			<TableBodyRow>
				<TableBodyCell class="!p-4">
					<Checkbox />
				</TableBodyCell>
	
			  <TableBodyCell>{item.Typ}</TableBodyCell>
			  <TableBodyCell>{item.BMK}</TableBodyCell>
			  <TableBodyCell>{item.Bearbeiter_SITECA}</TableBodyCell>
			  <TableBodyCell>{item.Bearbeiter_Kunde}</TableBodyCell>
			  <TableBodyCell>{item.Liefertermin}</TableBodyCell>
			  <TableBodyCell>{item.Prio}</TableBodyCell>
			</TableBodyRow>
			
		  {/each}
		</TableBody>
		<tfoot>
			<tr class="font-semibold text-gray-900 dark:text-white">
			  <th scope="row" class="py-3 px-6 text-base">Total</th>
			  <td class="py-3 px-6">3</td>
			  <td class="py-3 px-6">21,000</td>
			</tr>
		  </tfoot>
	</TabItem>
	<TabItem title="Arbeitsvorbereitung">
		<TableHead>
			<TableHeadCell class="!p-4">
				<Checkbox />
			</TableHeadCell>
		  <TableHeadCell>Typ</TableHeadCell>
		  <TableHeadCell>BMK/Einbauort</TableHeadCell>
		  <TableHeadCell>Bearbeiter SITECA</TableHeadCell>
		  <TableHeadCell>Bearbeiter Kunde</TableHeadCell>
		  <TableHeadCell>Liefertermin</TableHeadCell>
		  <TableHeadCell>Hinweis / Prio</TableHeadCell>
		  
		</TableHead>
		<TableBody tableBodyClass="divide-y">
		  {#each filteredItems as item}
			<TableBodyRow>
				<TableBodyCell class="!p-4">
					<Checkbox />
				</TableBodyCell>
	
			  <TableBodyCell>{item.Typ}</TableBodyCell>
			  <TableBodyCell>{item.BMK}</TableBodyCell>
			  <TableBodyCell>{item.Bearbeiter_SITECA}</TableBodyCell>
			  <TableBodyCell>{item.Bearbeiter_Kunde}</TableBodyCell>
			  <TableBodyCell>{item.Liefertermin}</TableBodyCell>
			  <TableBodyCell>{item.Prio}</TableBodyCell>
			</TableBodyRow>
			
		  {/each}
		</TableBody>
		<tfoot>
			<tr class="font-semibold text-gray-900 dark:text-white">
			  <th scope="row" class="py-3 px-6 text-base">Total</th>
			  <td class="py-3 px-6">3</td>
			  <td class="py-3 px-6">21,000</td>
			</tr>
		  </tfoot>
	</TabItem>
	<TabItem title="Fertigung mechanisch">
		<TableHead>
			<TableHeadCell class="!p-4">
				<Checkbox />
			</TableHeadCell>
		  <TableHeadCell>Typ</TableHeadCell>
		  <TableHeadCell>BMK/Einbauort</TableHeadCell>
		  <TableHeadCell>Bearbeiter SITECA</TableHeadCell>
		  <TableHeadCell>Bearbeiter Kunde</TableHeadCell>
		  <TableHeadCell>Liefertermin</TableHeadCell>
		  <TableHeadCell>Hinweis / Prio</TableHeadCell>
		  
		</TableHead>
		<TableBody tableBodyClass="divide-y">
		  {#each filteredItems as item}
			<TableBodyRow>
				<TableBodyCell class="!p-4">
					<Checkbox />
				</TableBodyCell>
	
			  <TableBodyCell>{item.Typ}</TableBodyCell>
			  <TableBodyCell>{item.BMK}</TableBodyCell>
			  <TableBodyCell>{item.Bearbeiter_SITECA}</TableBodyCell>
			  <TableBodyCell>{item.Bearbeiter_Kunde}</TableBodyCell>
			  <TableBodyCell>{item.Liefertermin}</TableBodyCell>
			  <TableBodyCell>{item.Prio}</TableBodyCell>
			</TableBodyRow>
			
		  {/each}
		</TableBody>
		<tfoot>
			<tr class="font-semibold text-gray-900 dark:text-white">
			  <th scope="row" class="py-3 px-6 text-base">Total</th>
			  <td class="py-3 px-6">3</td>
			  <td class="py-3 px-6">21,000</td>
			</tr>
		  </tfoot>
	</TabItem>
	<TabItem title="Fertigung elektrotechnisch">
		<TableHead>
			<TableHeadCell class="!p-4">
				<Checkbox />
			</TableHeadCell>
		  <TableHeadCell>Typ</TableHeadCell>
		  <TableHeadCell>BMK/Einbauort</TableHeadCell>
		  <TableHeadCell>Bearbeiter SITECA</TableHeadCell>
		  <TableHeadCell>Bearbeiter Kunde</TableHeadCell>
		  <TableHeadCell>Liefertermin</TableHeadCell>
		  <TableHeadCell>Hinweis / Prio</TableHeadCell>
		  
		</TableHead>
		<TableBody tableBodyClass="divide-y">
		  {#each filteredItems as item}
			<TableBodyRow>
				<TableBodyCell class="!p-4">
					<Checkbox />
				</TableBodyCell>
	
			  <TableBodyCell>{item.Typ}</TableBodyCell>
			  <TableBodyCell>{item.BMK}</TableBodyCell>
			  <TableBodyCell>{item.Bearbeiter_SITECA}</TableBodyCell>
			  <TableBodyCell>{item.Bearbeiter_Kunde}</TableBodyCell>
			  <TableBodyCell>{item.Liefertermin}</TableBodyCell>
			  <TableBodyCell>{item.Prio}</TableBodyCell>
			</TableBodyRow>
			
		  {/each}
		</TableBody>
		<tfoot>
			<tr class="font-semibold text-gray-900 dark:text-white">
			  <th scope="row" class="py-3 px-6 text-base">Total</th>
			  <td class="py-3 px-6">3</td>
			  <td class="py-3 px-6">21,000</td>
			</tr>
		  </tfoot>
	  </TabItem>
	  <TabItem title="Endmontage">
		<TableHead>
			<TableHeadCell class="!p-4">
				<Checkbox />
			</TableHeadCell>
		  <TableHeadCell>Typ</TableHeadCell>
		  <TableHeadCell>BMK/Einbauort</TableHeadCell>
		  <TableHeadCell>Bearbeiter SITECA</TableHeadCell>
		  <TableHeadCell>Bearbeiter Kunde</TableHeadCell>
		  <TableHeadCell>Liefertermin</TableHeadCell>
		  <TableHeadCell>Hinweis / Prio</TableHeadCell>
		  
		</TableHead>
		<TableBody tableBodyClass="divide-y">
		  {#each filteredItems as item}
			<TableBodyRow>
				<TableBodyCell class="!p-4">
					<Checkbox />
				</TableBodyCell>
	
			  <TableBodyCell>{item.Typ}</TableBodyCell>
			  <TableBodyCell>{item.BMK}</TableBodyCell>
			  <TableBodyCell>{item.Bearbeiter_SITECA}</TableBodyCell>
			  <TableBodyCell>{item.Bearbeiter_Kunde}</TableBodyCell>
			  <TableBodyCell>{item.Liefertermin}</TableBodyCell>
			  <TableBodyCell>{item.Prio}</TableBodyCell>
			</TableBodyRow>
			
		  {/each}
		</TableBody>
		<tfoot>
			<tr class="font-semibold text-gray-900 dark:text-white">
			  <th scope="row" class="py-3 px-6 text-base">Total</th>
			  <td class="py-3 px-6">3</td>
			  <td class="py-3 px-6">21,000</td>
			</tr>
		  </tfoot>
	  </TabItem>
	  <TabItem title="Administration">
		<TableHead>
			<TableHeadCell class="!p-4">
				<Checkbox />
			</TableHeadCell>
		  <TableHeadCell>Typ</TableHeadCell>
		  <TableHeadCell>BMK/Einbauort</TableHeadCell>
		  <TableHeadCell>Bearbeiter SITECA</TableHeadCell>
		  <TableHeadCell>Bearbeiter Kunde</TableHeadCell>
		  <TableHeadCell>Liefertermin</TableHeadCell>
		  <TableHeadCell>Hinweis / Prio</TableHeadCell>
		  
		</TableHead>
		<TableBody tableBodyClass="divide-y">
		  {#each filteredItems as item}
			<TableBodyRow>
				<TableBodyCell class="!p-4">
					<Checkbox />
				</TableBodyCell>
	
			  <TableBodyCell>{item.Typ}</TableBodyCell>
			  <TableBodyCell>{item.BMK}</TableBodyCell>
			  <TableBodyCell>{item.Bearbeiter_SITECA}</TableBodyCell>
			  <TableBodyCell>{item.Bearbeiter_Kunde}</TableBodyCell>
			  <TableBodyCell>{item.Liefertermin}</TableBodyCell>
			  <TableBodyCell>{item.Prio}</TableBodyCell>
			</TableBodyRow>
			
		  {/each}
		</TableBody>
		<tfoot>
			<tr class="font-semibold text-gray-900 dark:text-white">
			  <th scope="row" class="py-3 px-6 text-base">Total</th>
			  <td class="py-3 px-6">3</td>
			  <td class="py-3 px-6">21,000</td>
			</tr>
		  </tfoot>
	  </TabItem>
	  <TabItem title="Reserve">
		<TableHead>
			<TableHeadCell class="!p-4">
				<Checkbox />
			</TableHeadCell>
		  <TableHeadCell>Typ</TableHeadCell>
		  <TableHeadCell>BMK/Einbauort</TableHeadCell>
		  <TableHeadCell>Bearbeiter SITECA</TableHeadCell>
		  <TableHeadCell>Bearbeiter Kunde</TableHeadCell>
		  <TableHeadCell>Liefertermin</TableHeadCell>
		  <TableHeadCell>Hinweis / Prio</TableHeadCell>
		  
		</TableHead>
		<TableBody tableBodyClass="divide-y">
		  {#each filteredItems as item}
			<TableBodyRow>
				<TableBodyCell class="!p-4">
					<Checkbox />
				</TableBodyCell>
	
			  <TableBodyCell>{item.Typ}</TableBodyCell>
			  <TableBodyCell>{item.BMK}</TableBodyCell>
			  <TableBodyCell>{item.Bearbeiter_SITECA}</TableBodyCell>
			  <TableBodyCell>{item.Bearbeiter_Kunde}</TableBodyCell>
			  <TableBodyCell>{item.Liefertermin}</TableBodyCell>
			  <TableBodyCell>{item.Prio}</TableBodyCell>
			</TableBodyRow>
			
		  {/each}
		</TableBody>
		<tfoot>
			<tr class="font-semibold text-gray-900 dark:text-white">
			  <th scope="row" class="py-3 px-6 text-base">Total</th>
			  <td class="py-3 px-6">3</td>
			  <td class="py-3 px-6">21,000</td>
			</tr>
		  </tfoot>
	  </TabItem>
  </Tabs>
</TableSearch>


<style>
	section {
		display: flex;
		flex-direction: column;
		justify-content: center;
		align-items: center;
		flex: 0.6;
	}



	.box {
		padding: 5rem;
		height: 100%;
	  width: 100%;
	  top: 100px;
	  position: absolute;
	}
</style>
