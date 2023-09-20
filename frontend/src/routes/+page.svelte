<script>
    
    import { getAllContexts, onMount } from 'svelte';

    const baseUrl = "http://localhost:8080"
    
    let immoScoutRegions = [];

    let whatSearch = "";
    let whereSearchHelper;
    let whereSearch = [];
    let priceStart;
    let roomSquareStart;
    let roomNumberHelper;
    let roomNumber;
    let whatKind;

    

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
        //console.log(whatSearch, whereSearch, priceStart)
        whereSearchHelper = whereSearchHelper.replace(/ / , "")
        whereSearchHelper = whereSearchHelper.replace(/ / , "")
        whereSearchHelper = whereSearchHelper.replace(/ / , "")
        
        whereSearch = whereSearchHelper.split(",")
        whereSearch.splice(3, 1) // schmeisst falsches raus 

        if (roomNumberHelper == "ab 1 Zimmer") {
            roomNumber = 1.0;
        } else if (roomNumberHelper == "ab 1,5 Zimmer") {
            roomNumber = 1.5;
        } else if (roomNumberHelper == "ab 2 Zimmer") {
            roomNumber = 2.0;
        } else if (roomNumberHelper == "ab 2,5 Zimmer") {
            roomNumber = 2.5;
        } else if (roomNumberHelper == "ab 3 Zimmer") {
            roomNumber = 3.0;
        } else if (roomNumberHelper == "ab 3,5 Zimmer") {
            roomNumber = 3.5;
        } else if (roomNumberHelper == "ab 4 Zimmer") {
            roomNumber = 4.0;
        } else if (roomNumberHelper == "ab 5 Zimmer") {
            roomNumber = 5.0;
        } else if (roomNumberHelper == "Zimmer egal") {
            roomNumber = 0;
        }

        console.log(roomNumber)

        console.log(whereSearchHelper)
        const params = {
            "What":  whatSearch ,
            "Kind":  whatKind,    
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
                <label for="browser">Art?</label>
                <input list="houseKind" name="whatKindOfHouse" id="whatKindOfHouse" placeholder="kaufen" bind:value={ whatKind }>
                    <datalist id="houseKind">
                        <option value="kaufen">
                        <option value="mieten">
                        <option value="bauen">
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

        <p>Preis ab</p><input type="number" min="0" bind:value={ priceStart }>

        <p>Wohnfläche ab</p><input type="number" min="0" bind:value={ roomSquareStart }>
    
    
        

        <div class="searchOption">
            <form>
                <!---->
                <label for="browser">Zimmer</label>
                <input list="numberOfRooms" name="roomNumber" id="roomNumber" placeholder="Zimmer ab" bind:value= { roomNumberHelper }>
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