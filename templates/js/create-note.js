
async function createNote() {
    var title = document.querySelector("#input-title").value;
    var element = document.querySelector(".trix-custom-content");
    // var document = element.editor.getDocument();
    console.log(title);
    // console.log(element);
    //need to extract data from trix content and store
    //need to extract attachments and save via AWS s3 bucket
    const data = {
        title: title,
        body: "<div> Dummy content </div>",
    }
    const headers = {
        'Content-Type': 'application/json'
    }
    await axios.post(ROOT_API + "api/notes/", data, headers)
    .then(response => {
        console.log(response);
        if (response.status == 201) {
            window.location.href = ROOT_API + "view-notes";
        }
    });
}