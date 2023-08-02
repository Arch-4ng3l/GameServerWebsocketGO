<script>
    import { writable } from "svelte/store"
    let object = {
        name: "Test",
        x: 0,
        y: 0,
        z: 0,
    };
    let player = {
        name: "Player",
        x: 0,
        y: 0,
        z: 0,
    };
    let objects = writable([]);
    let socket = new WebSocket("ws://localhost:3000/api/");

    let jsonPlayer = JSON.stringify(player);
    socket.onmessage = (event) => {
        console.log(event.data);
        let obj = JSON.parse(event.data);
        objects.set([]);
        obj.forEach(element => {
            if (element.name !== "") {
                objects.update(arr => [...arr, element])
            }
        }
        );
    }
    function fetchObjects() {
        socket.send(jsonPlayer);
    }

    function sendObject() {
        let name = document.getElementById("name").value;
        let x = document.getElementById("x").value;
        let y = document.getElementById("y").value;
        let z = document.getElementById("z").value;
        if (name === "") {
            return
        }
        console.log(name);
        object.name = name;
        object.x = parseFloat(x);
        object.y = parseFloat(y);
        object.z = parseFloat(z);
        let jsonObj = JSON.stringify(object);
        socket.send(jsonObj); 
        fetchObjects();
    }

</script>


<div>
    <label for="">Name: </label>
    <input name="name" type="text" id="name"/>

    <label for="x">X: </label>
    <input name="x" type="number" id="x" value="0"/>

    <label for="y">Y: </label>
    <input name="y" type="number" id="y" value="0"/>

    <label for="z">Z: </label>
    <input name="z" type="number" id="z" value="0"/>
    <br>
    <button on:click={sendObject} value="send">send</button>
    <button on:click={fetchObjects} value="fetch">fetch</button>
    <div class="list">
    {#each $objects as {name, x, y, z}}
        <div class="obj">
            <h2>{name}</h2>
            <p>x : {x}</p>
            <p>y : {y}</p>
            <p>z : {z}</p>
        </div>
        <br>
        <br>
    {/each}
    </div>
</div>

<style>
    .obj {
        text-align: center;
        padding: 10px;
        width: 5em;
        outline: 3px;
        border-radius: 10px;
        border-style:solid;
        border-color:black;
        margin: 10px;
    }
    .list {
        padding: 10px;
        display: flex;
        justify-content: center;
        right: 1em;
        align-content: center;
    }
</style>
