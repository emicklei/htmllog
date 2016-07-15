package htmllog

var DefaultScrollToBottom = `<script src="http://ajax.googleapis.com/ajax/libs/jquery/1.11.3/jquery.min.js"></script>
<script type="text/javascript">
 	function loaded() {
	  jQuery('html body').animate(
		  { scrollTop: $(document).height() }, 
		  50,
		  "linear");
	}
</script>`
