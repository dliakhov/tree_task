package tree_test

import (
	"testing"

	"github.com/dliakhov/tree_task/tree"
)

func TestIsSymetric(t *testing.T) {
	tests := []struct {
		name string
		tree *tree.Tree
		want bool
	}{
		//        5
		//       / \
		//      7   3
		{
			name: "should return is not symetric 2 children nodes",
			tree: &tree.Tree{
				Root: &tree.Node{
					Val: 5,
					Left: &tree.Node{
						Val: 7,
					},
					Right: &tree.Node{
						Val: 3,
					},
				},
			},
			want: false,
		},
		//        5
		//       / \
		//      7   7
		{
			name: "should return is symetric 2 children nodes",
			tree: &tree.Tree{
				Root: &tree.Node{
					Val: 5,
					Left: &tree.Node{
						Val: 7,
					},
					Right: &tree.Node{
						Val: 7,
					},
				},
			},
			want: true,
		},
		//        5
		//       / \
		//      7   7
		//     / \ / \
		//       3 3
		{
			name: "should return is symetric 4 children nodes",
			tree: &tree.Tree{
				Root: &tree.Node{
					Val: 5,
					Left: &tree.Node{
						Val: 7,
						Right: &tree.Node{
							Val: 3,
						},
					},
					Right: &tree.Node{
						Val: 7,
						Left: &tree.Node{
							Val: 3,
						},
					},
				},
			},
			want: true,
		},
		//        5
		//       / \
		//      7   7
		//     / \ / \
		//       3
		{
			name: "should return is symetric 3 children nodes",
			tree: &tree.Tree{
				Root: &tree.Node{
					Val: 5,
					Left: &tree.Node{
						Val: 7,
						Right: &tree.Node{
							Val: 3,
						},
					},
					Right: &tree.Node{
						Val: 7,
					},
				},
			},
			want: false,
		},
		//        5
		//       / \
		//      7   7
		//     / \ / \
		//    3  4 4  3
		//   / \     / \
		//  1   6   6   1
		{
			name: "should return is symetric 10 children nodes",
			tree: &tree.Tree{
				Root: &tree.Node{
					Val: 5,
					Left: &tree.Node{
						Val: 7,
						Left: &tree.Node{
							Val: 3,
							Left: &tree.Node{
								Val: 1,
							},
							Right: &tree.Node{
								Val: 6,
							},
						},
						Right: &tree.Node{
							Val: 4,
						},
					},
					Right: &tree.Node{
						Val: 7,
						Left: &tree.Node{
							Val: 4,
						},
						Right: &tree.Node{
							Val: 3,
							Left: &tree.Node{
								Val: 6,
							},
							Right: &tree.Node{
								Val: 1,
							},
						},
					},
				},
			},
			want: true,
		},

		//        5
		//       / \
		//      7   7
		//     / \ / \
		//    3  4 4  3
		//   / \     / \
		//  1   6   6   1
		// / \ /   / \ / \
		// 0 2 8     8 2  0
		{
			name: "should return is symetric 16 children nodes",
			tree: &tree.Tree{
				Root: &tree.Node{
					Val: 5,
					Left: &tree.Node{
						Val: 7,
						Left: &tree.Node{
							Val: 3,
							Left: &tree.Node{
								Val: 1,
								Left: &tree.Node{
									Val: 0,
								},
								Right: &tree.Node{
									Val: 2,
								},
							},
							Right: &tree.Node{
								Val: 6,
								Left: &tree.Node{
									Val: 8,
								},
							},
						},
						Right: &tree.Node{
							Val: 4,
						},
					},
					Right: &tree.Node{
						Val: 7,
						Left: &tree.Node{
							Val: 4,
						},
						Right: &tree.Node{
							Val: 3,
							Left: &tree.Node{
								Val: 6,
								Right: &tree.Node{
									Val: 8,
								},
							},
							Right: &tree.Node{
								Val: 1,
								Left: &tree.Node{
									Val: 2,
								},
								Right: &tree.Node{
									Val: 0,
								},
							},
						},
					},
				},
			},
			want: true,
		},
		//        5
		//       / \
		//      7   7
		//     / \ / \
		//    3  4 4  3
		//   / \     / \
		//  1   6   6   1
		// / \ /   / \ / \
		// 0 2 8   8  8 2  0
		{
			name: "should return is not symetric 17 children nodes",
			tree: &tree.Tree{
				Root: &tree.Node{
					Val: 5,
					Left: &tree.Node{
						Val: 7,
						Left: &tree.Node{
							Val: 3,
							Left: &tree.Node{
								Val: 1,
								Left: &tree.Node{
									Val: 0,
								},
								Right: &tree.Node{
									Val: 2,
								},
							},
							Right: &tree.Node{
								Val: 6,
								Left: &tree.Node{
									Val: 8,
								},
								Right: &tree.Node{
									Val: 8,
								},
							},
						},
						Right: &tree.Node{
							Val: 4,
						},
					},
					Right: &tree.Node{
						Val: 7,
						Left: &tree.Node{
							Val: 4,
						},
						Right: &tree.Node{
							Val: 3,
							Left: &tree.Node{
								Val: 6,
								Right: &tree.Node{
									Val: 8,
								},
							},
							Right: &tree.Node{
								Val: 1,
								Left: &tree.Node{
									Val: 2,
								},
								Right: &tree.Node{
									Val: 0,
								},
							},
						},
					},
				},
			},
			want: false,
		},
		//          5
		//         / \
		//        7   7
		//       / \ / \
		//      3  4 4  3
		//     / \     / \
		//    1   6   6   1
		//   / \ /   / \ / \
		//   0 2 8     8 2  0
		//  /                \
		// 5                  5
		{
			name: "should return is not symetric 18 children nodes",
			tree: &tree.Tree{
				Root: &tree.Node{
					Val: 5,
					Left: &tree.Node{
						Val: 7,
						Left: &tree.Node{
							Val: 3,
							Left: &tree.Node{
								Val: 1,
								Left: &tree.Node{
									Val: 0,
									Left: &tree.Node{
										Val: 5,
									},
								},
								Right: &tree.Node{
									Val: 2,
								},
							},
							Right: &tree.Node{
								Val: 6,
								Left: &tree.Node{
									Val: 8,
								},
							},
						},
						Right: &tree.Node{
							Val: 4,
						},
					},
					Right: &tree.Node{
						Val: 7,
						Left: &tree.Node{
							Val: 4,
						},
						Right: &tree.Node{
							Val: 3,
							Left: &tree.Node{
								Val: 6,
								Right: &tree.Node{
									Val: 8,
								},
							},
							Right: &tree.Node{
								Val: 1,
								Left: &tree.Node{
									Val: 2,
								},
								Right: &tree.Node{
									Val: 0,
									Right: &tree.Node{
										Val: 5,
									},
								},
							},
						},
					},
				},
			},
			want: true,
		},
		//        5
		//       / \
		//      7   7
		//     / \ / \
		//    3  4 4  3
		//   / \     / \
		//  1   6   6
		{
			name: "should return is symetric 9 children nodes",
			tree: &tree.Tree{
				Root: &tree.Node{
					Val: 5,
					Left: &tree.Node{
						Val: 7,
						Left: &tree.Node{
							Val: 3,
							Left: &tree.Node{
								Val: 1,
							},
							Right: &tree.Node{
								Val: 6,
							},
						},
						Right: &tree.Node{
							Val: 4,
						},
					},
					Right: &tree.Node{
						Val: 7,
						Left: &tree.Node{
							Val: 4,
						},
						Right: &tree.Node{
							Val: 3,
							Left: &tree.Node{
								Val: 6,
							},
						},
					},
				},
			},
			want: false,
		},
		//        5
		//       / \
		//      7   7
		//     / \ / \
		//    3  4 4  3
		//   / \     / \
		//  1   6   1   1
		{
			name: "should return is not symetric 10 children nodes",
			tree: &tree.Tree{
				Root: &tree.Node{
					Val: 5,
					Left: &tree.Node{
						Val: 7,
						Left: &tree.Node{
							Val: 3,
							Left: &tree.Node{
								Val: 1,
							},
							Right: &tree.Node{
								Val: 6,
							},
						},
						Right: &tree.Node{
							Val: 4,
						},
					},
					Right: &tree.Node{
						Val: 7,
						Left: &tree.Node{
							Val: 4,
						},
						Right: &tree.Node{
							Val: 3,
							Left: &tree.Node{
								Val: 1,
							},
							Right: &tree.Node{
								Val: 1,
							},
						},
					},
				},
			},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := tree.IsSymetric(tt.tree)
			if res != tt.want {
				t.Errorf("got %v, want %v", res, tt.want)
			}
		})
	}
}
