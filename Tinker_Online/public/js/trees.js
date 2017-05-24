var treeData = [];
var i = 0,
duration = 350,
root, tree, diagonal, svg;
  var margin = {top: 40, right: 120, bottom: 20, left: 120},
    width = 960 - margin.right - margin.left,
    height = 500 - margin.top - margin.bottom;

function initTree() {
  // ************** Generate the tree diagram	 *****************
  tree = d3.layout.tree()
  	.size([height, width]);
  diagonal = d3.svg.diagonal()
  	.projection(function(d) { return [d.x, d.y]; });
  svg = d3.select("#tree").append("svg")
  	.attr("width", width + margin.right + margin.left)
  	.attr("height", height + margin.top + margin.bottom)
    .append("g")
  	.attr("transform", "translate(" + margin.left + "," + margin.top + ")");
};

function updateTree() {
   if (treeData != []) {
    root = treeData;
    root.x0 = height / 2;
    root.y0 = 0;
    update(root);
    d3.select(self.frameElement).style("height", "500px");
  };
}

function update(source) {
  // Compute the new tree layout.
  var nodes = tree.nodes(root).reverse(),
	  links = tree.links(nodes);
  // Normalize for fixed-depth.
  nodes.forEach(function(d) { d.y = d.depth * 100; });
  // Update the nodes…
  var node = svg.selectAll("g.node")
	  .data(nodes, function(d) { return d.id || (d.id = ++i); });
  // Enter any new nodes at the parent's previous position.
  var nodeEnter = node.enter().append("g")
	  .attr("class", "node")
	  .attr("transform", function(d) { return "translate(" + source.x0 + "," + source.y0 + ")"; })
	  .on("click", click);
  nodeEnter.append("circle")
	  .attr("r", 1e-6)
	  .style("fill", function(d) { return d._children ? "lightsteelblue" : "#fff"; });
  nodeEnter.append("text")
	  .attr("y", function(d) { return d.children || d._children ? -18 : 18; })
	  .attr("dy", ".35em")
	  .attr("text-anchor", "middle")
	  .text(function(d) { return d.value; })
	  .style("fill-opacity", 1);
  // Transition nodes to their new position.
  var nodeUpdate = node.transition()
	  .duration(duration)
	  .attr("transform", function(d) { return "translate(" + d.x + "," + d.y + ")"; });
  nodeUpdate.select("circle")
	  .attr("r", 10)
	  .style("fill", function(d) { return d._children ? "lightsteelblue" : "#fff"; });
  nodeUpdate.select("text")
	  .style("fill-opacity", 1);
  // Transition exiting nodes to the parent's new position.
  var nodeExit = node.exit().transition()
	  .duration(duration)
	  .attr("transform", function(d) { return "translate(" + source.y + "," + source.x + ")"; })
	  .remove();
  nodeExit.select("circle")
	  .attr("r", 1e-6);
  nodeExit.select("text")
	  .style("fill-opacity", 1e-6);
  // Update the links…
  var link = svg.selectAll("path.link")
	  .data(links, function(d) { return d.target.id; });
  // Enter any new links at the parent's previous position.
  link.enter().insert("path", "g")
	  .attr("class", "link")
	  .attr("d", function(d) {
		var o = {x: source.x0, y: source.y0};
		return diagonal({source: o, target: o});
	  });
  // Transition links to their new position.
  link.transition()
	  .duration(duration)
	  .attr("d", diagonal);
  // Transition exiting nodes to the parent's new position.
  link.exit().transition()
	  .duration(duration)
	  .attr("d", function(d) {
		var o = {x: source.x, y: source.y};
		return diagonal({source: o, target: o});
	  })
	  .remove();
  // Stash the old positions for transition.
  nodes.forEach(function(d) {
	d.x0 = d.x;
	d.y0 = d.y;
  });
}
// Toggle children on click.
var sm="";

function evalChildren(d) {
  if (!d.children && !d._children) {
    sm+=d.value;
  }
  else if (d._children) {
    evalChildren(d._children[0])
    sm=sm+d.value
    evalChildren(d._children[1])
  }
  else {
    evalChildren(d.children[0])
    sm=sm+d.value;
    evalChildren(d.children[1])
  }
}

function repPower() {
  for(var i=0;i<sm.length-1;i++) {
    console.log(i);
    sm=sm.replace("^","**");
  }
}

function click(d) {
  if (d.children) {
      sm="";
      evalChildren(d);
      repPower();
      alert(sm + "=" + eval(sm));
	     d._children = d.children;
	     d.children = null;
  } else {
	     d.children = d._children;
	     d._children = null;
  }
    update(d);
}