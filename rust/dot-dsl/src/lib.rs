pub mod graph {
    pub mod graph_items {
        pub mod node {
            use std::collections::HashMap;

            #[derive(Clone, Debug, PartialEq)]
            pub struct Node {
                pub name: String,
                attrs: HashMap<String, String>,
            }

            impl Node {
                pub fn new(name: &str) -> Self {
                    Node { name: name.to_string(), attrs: HashMap::default() }
                }

                pub fn with_attrs(mut self, attrs: &[(&str, &str)]) -> Self {
                    self.attrs.extend(
                        attrs
                            .iter()
                            .cloned()
                            .map(|(name, value)| (name.to_owned(), value.to_owned()))
                    );

                    self
                }

                pub fn get_attr<'a>(&'a self, name: &str) -> Option<&'a str> {
                    self.attrs.get(name).map(|value| value.as_str())
                }
            }
        }

        pub mod edge {
            use std::collections::HashMap;

            #[derive(Clone, Debug, PartialEq)]
            pub struct Edge {
                from: String,
                to: String,
                attrs: HashMap<String, String>,
            }

            impl Edge {
                pub fn new(from: &str, to: &str) -> Self {
                    Edge { from: from.to_string(), to: to.to_string(), attrs: HashMap::default() }
                }

                pub fn with_attrs(mut self, attrs: &[(&str, &str)]) -> Self {
                    self.attrs.extend(
                        attrs
                            .iter()
                            .cloned()
                            .map(|(name, value)| (name.to_owned(), value.to_owned()))
                    );

                    self
                }
            }
        }
    }

    use std::collections::HashMap;

    use graph_items::{node::Node, edge::Edge};

    pub struct Graph {
        pub nodes: Vec<Node>,
        pub edges: Vec<Edge>,
        pub attrs: HashMap<String, String>,
    }

    impl Graph {
        pub fn new() -> Self {
            Graph { nodes: Vec::default(), edges: Vec::default(), attrs: HashMap::default() }
        }

        pub fn with_nodes(mut self, nodes: &Vec<Node>) -> Self {
            self.nodes.extend(nodes.iter().cloned());
            self
        }

        pub fn with_edges(mut self, edges: &Vec<Edge>) -> Self {
            self.edges.extend(edges.iter().cloned());
            self
        }

        pub fn with_attrs(mut self, attrs: &[(&str, &str)]) -> Self {
            self.attrs.extend(
            attrs
                    .iter()
                    .cloned()
                    .map(|(name, value)| (name.to_owned(), value.to_owned()))
            );

            self
        }

        pub fn get_node<'a>(&'a self, name: &str) -> Option<&'a Node> {
            self.nodes.iter().find(|node| node.name == name)
        }
    }
}
