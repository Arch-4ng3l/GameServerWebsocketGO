
<script>
    let fileIn;
    function handleSendFile() {
        const file = fileIn.files[0];
        const url = "/api/assets";
        const formData = new FormData();
        formData.append("file", file);


        fetch(url, {
            method: "POST",
            body: formData,
        })
        .then(response => {
            if (response.ok) {
                console.log("success");
            }
        });

    }
</script>

<div>
    <h2>Assets</h2>
    <br>
    <br>
    <input  class="in" bind:this={fileIn} type="file" id="fileInput"/>
    <br>
    <br>
    <button class="btn" on:click={handleSendFile}>Upload</button> 
</div>

<style>
    
.in {
    border-radius: 10px;
    width: 40%;
}
.in:hover {
    border-color: #ff3e00;
}

.btn {
    border-radius: 10px;
    transition: ease-in-out 0.2s;
}
.btn:hover {
    border-color: #ff3e00;
    color : #ff3e00;
    font-weight: bold;
}
</style>
