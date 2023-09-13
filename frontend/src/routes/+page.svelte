<script>
    
    import { getAllContexts, onMount } from 'svelte';

    const baseUrl = "http://localhost:8080"
    
    let immoScoutRegions = [];

    let whatSearch = "";
    let whereSearchHelper;
    let whereSearch = [];
    let priceStart;
    let roomSquareStart;
    let roomNumber;

    

    onMount(async () => {
		getImmoScoutRegions();
	});

    async function getImmoScoutRegions() {
        try {
            let res = await fetch(baseUrl + "/immoScoutRegions");
            console.log(res)
            immoScoutRegions = await res.json();
            console.log(immoScoutRegions);
        } catch (error) {
            err = true;
        }
    }

    async function onClickedSearch() {
        console.log(whatSearch, whereSearch, priceStart)
        whereSearchHelper = whereSearchHelper.replace(/ / , "")
        whereSearchHelper = whereSearchHelper.replace(/ / , "")
        whereSearchHelper = whereSearchHelper.replace(/ / , "")
        
        whereSearch = whereSearchHelper.split(",")
        whereSearch = whereSearch.splice(1, 3)
        console.log(whereSearchHelper)
        const params = {
            "What":  whatSearch ,
            "Where":  whereSearch ,
            "Price":  priceStart ,
            "Squaremeters":  roomSquareStart ,
            "RoomNumber" :  parseFloat(roomNumber),
        };
        console.log(params);
        const response = await fetch(baseUrl + "/search", {
            method: 'POST',
            headers: {
                'Accept': 'application/json',
                'Content-Type': 'application/json'
            },
            body: JSON.stringify(params),
        });
        const results = await response.json();
    }

</script>



<header>
    <div id="headline">
        <h1>Immo-Scraper</h1>
        <h3>Welcome to Immo-Scraper, your Real Estate Scraper for Germany</h3>
    </div>

    <div id="searchOptions">
        <div class="searchOption">
            <form>
                <!---->
                <label for="browser">Was?</label>
                <input list="houseTypes" name="whatTypeOfHouse" id="whatTypeOfHouse" placeholder="Haus" bind:value={ whatSearch }>
                    <datalist id="houseTypes">
                        <option value="Haus">
                        <option value="Wohnung">
                        <option value="Wohnen auf Zeit">
                        <option value="WG-Zimmer">
                        <option value="Büros/Praxen">
                        <option value="Hallen/Produktion">
                        <option value="Einzelhandel">
                        <option value="Lagerräume/-boxen">
                        <option value="Spezialgewerbe">
                        <option value="Grundstücke">
                        <option value="Gewerbegrundstücke">
                        <option value="Garagen/Stellplätze">
                        <option value="Serioenwohnheim">
                        <option value="Pflegeheim">
                    </datalist>
            </form>
        </div>

        <div class="searchOption">
            <form>
                <!---->
                <label for="browser">Wo?</label>
                <input list="whereLocated" name="location" id="location" placeholder="Stadt" bind:value= { whereSearchHelper }>
                    <datalist id="whereLocated">
                        { #each immoScoutRegions as { bundesland, kreis, stadt } } 
                        <option value= "{ stadt}, { bundesland } { kreis }" />
                        { /each }
                    </datalist>
            </form>
        </div>

        <p>Preis ab</p><input type="number" bind:value={ priceStart }>

        <p>Wohnfläche ab</p><input type="number" bind:value={ roomSquareStart }>
    
    
        

        <div class="searchOption">
            <form>
                <!---->
                <label for="browser">Zimmer</label>
                <input list="numberOfRooms" name="roomNumber" id="roomNumber" placeholder="Zimmer ab" bind:value= { roomNumber }>
                    <datalist id="numberOfRooms">
                        <option value="Zimmer egal"></option>
                        <option value="ab 1 Zimmer"></option>
                        <option value="ab 1,5 Zimmer"></option>
                        <option value="ab 2 Zimmer"></option>
                        <option value="ab 2,5 Zimmer"></option>
                        <option value="ab 3 Zimmer"></option>
                        <option value="ab 3,5 Zimmer"></option>
                        <option value="ab 4 Zimmer"></option>
                        <option value="ab 5 Zimmer"></option>
                    </datalist>
            </form>
        </div>

    </div> 


    <button type="submit" on:click= { () => onClickedSearch() } >Suchen</button>

</header>






<style>
    #location{
        width: 50rem;
    }
</style>