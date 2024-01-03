
//method will be applied to edit-note.js and applicable under Create note, Clone note, Edit note
const checkFields = () => {
    const inputTitle = document.getElementById('input-title');
    const formContent = document.getElementById('form-content');
    const publishBtn = document.querySelector('.publish-btn');
    const titleValue = inputTitle.value.trim();
    const contentValue = formContent.value.trim();

    if (titleValue === '' || contentValue === '') {
        publishBtn.disabled = true;
    } else {
        publishBtn.disabled = false;
    }
};

window.addEventListener('load', () => {
    const publishBtn = document.querySelector('.publish-btn');
    publishBtn.disabled = true;
});


async function createNote() {
    var title = document.querySelector("#input-title").value;
    const formContent = document.getElementById('form-content');
    const contentValue = formContent.value.trim();
    // var trix = document.querySelector(".trix-custom-content");
    // console.log(trix);
    // const editorContent = trix.editor.getDocument().toString();
    // console.log(trix.editor.getDocument());
    // const tempDiv = document.createElement('div');
    // tempDiv.innerHTML = editorContent;
    // const images = tempDiv.querySelectorAll('img');
    // console.log(images);
    // const imageUrls = [];
    // images.forEach((img) => {
    //     imageUrls.push(img.src);
    // });
    // console.log(imageUrls);

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