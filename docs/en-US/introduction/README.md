---
name: Introduction
---

_**ASoulDocs**_ is a web server for multilingual, real-time synchronized and searchable documentation.

## Motivation

Project documentation tools is an already-crowded place yet not being both mature and affordable for individual especially OSS developers. Countless static site generators, documentation servers, SaaS products, unfortunately, we are not happy with any of them. More importantly, we love and are capable of hacking on this area.

It has been years we struggled with picking one of them, that's basically why _**ASoulDocs**_ was created (previously "Peach" pre-1.0).

The following table illustrates the features (that we care) comparisons between _**ASoulDocs**_ and other existing tools (that we investigated, concepts may be different from what we understand):

|Name/Feature            |_**ASoulDocs**_|[Mkdocs](https://www.mkdocs.org/)|[Hugo](https://gohugo.io/)|[VuePress](https://v2.vuepress.vuejs.org/)/[VitePress](https://vitepress.vuejs.org/)|[GitBook](https://www.gitbook.com/)|[Notion](https://www.notion.so/)|
|:----------------------:|:-------------:|:----:|:--:|:----------------:|:----:|:-----:|
|Self-hosted             | ✅ | ✅ | ✅ | ✅ | ❌ | ❌ |
|Multilingual<sup>1</sup>| ✅ | ✅ | ✅ | ✅ | ❌ | ❌ |
|Builtin push-to-sync    | ✅ | ❌ | ❌ | ❌ | ✅ | ✅ |
|Algolia search          | 🎯 | ❌ | ✅ | ✅ | ❌ | ❌ |
|Commenting system       | 🎯 | ❌ | ✅ | ❌ | ❌ | ✅ |
|Versionable             | 🎯 | ❌ | ❌ | ❌ | ❌ | ❌ |
|Protected resources     | 🎯 | ❌ | ❌ | ❌ | ❌ | ❌ |
|Dark mode               | ✅ | ❌ | ✅ | ✅ | ❌ | ✅ |
|Customizable<sup>2</sup>| ✅ | ❌ | ✅ | ❌ | ❌ | ❌ |

- <sup>1</sup>: None of others support multilingual without changing the URL, which to us is a bizarre because we have to share different URLs for different groups of users.
- <sup>2</sup>: In such way that visitors couldn't recognize what is powering the site behind the scene.
- 🎯: Features that are on the roadmap.

## History

The project was initially named as "Peach Docs" in pre-1.0 releases, which is now branded as _**ASoulDocs**_ starting from the 1.0 release.

The tech stack has evolved since 2015, [Macaron](https://go-macaron.com) and [Semantic UI](https://semantic-ui.com/) was the new and hot things, and the latest golden partners are [Flamego](https://flamego.dev) and [Tailwind CSS](https://tailwindcss.com/).

The project is now part of [A-SOUL SIG](https://github.com/asoul-sig) (previously "github.com/peachdocs"), consists a group of A-SOUL fans.

## Getting started

_Stay tuned!_
