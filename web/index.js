function handleCopy(output){
    if (output == ""){
       alert("Nothing to copy. Please generate your art first.")
    }else{
        navigator.clipboard.writeText(output)
        alert("Art Successfully copied to clipboard")
    }
}

function handleDownload(output){
    if (output == ""){
       alert("No content available to download. Please submit the form.")
    }else{
        downloadFile(String(output),"art.txt","text/plain")
    }
}

function downloadFile(content, fileName, contentType) {
    const a = document.createElement("a");
    const file = new Blob([content], { type: contentType });
    a.href = URL.createObjectURL(file);
    a.download = fileName;
    a.click();
    URL.revokeObjectURL(a.href); // Clean up memory
}
