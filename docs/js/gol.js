 $(".button-collapse").sideNav();

 /*

  <button id="myBtn">Open Modal</button>

<div id="codeModal" class="modal">
  <div class="modal-content">
    <div class="modal-header">
      <span class="close">&times;</span>
      <h2>Modal Header</h2>
    </div>
    <div class="modal-body">
      <p>Some text in the Modal Body</p>
      <p>Some other text...</p>
    </div>
    <div class="modal-footer">
      <a href="">read code in lib here</a>
    </div>
  </div>
</div>

*/

function showCode(code_details){
  var modal = document.getElementById('codeModal');
  modal.style.display = "block";

  var btn = document.getElementById("myBtn");

  var span = document.getElementsByClassName("close")[0];

  span.onclick = function() {
    modal.style.display = "none";
  }

  window.onclick = function(event) {
    if (event.target == modal) {
      modal.style.display = "none";
    }
  }
}
