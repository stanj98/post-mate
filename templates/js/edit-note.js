
window.addEventListener('load', () => {
    const publishBtn = document.querySelector('.publish-btn');
    publishBtn.disabled = false;
});


async function editNote() {
    var noteID = document.getElementById("noteID").value;
    var title = document.querySelector("#input-title").value;
    const formContent = document.getElementById('form-content');
    const contentValue = formContent.value.trim();
    // var element = document.querySelector(".trix-custom-content");
    // var document = element.editor.getDocument();
    if (title === '' || contentValue === '') {
        alert("Please add a title and body to your note!");
        return;
    }
    
    //need to extract data from trix content and store
    //need to extract attachments and save via AWS s3 bucket
    const data = {
        title: title,
        body: "<div> Dummy content </div>",
    }
    await axios.put(ROOT_API + "api/notes/" + noteID, data)
    .then(response => {
        console.log(response);
        if (response.status == 200) {
            window.location.href = ROOT_API + noteID;
        }
    });
}